import { useState, useEffect, useMemo } from "react";
import Layout from "../components/Layout";
import InputPopup from "../components/InputPopup";

export default function Leaderboard() {
  const [models, setModels] = useState([]);
  const [selectedModel, setSelectedModel] = useState(null);
  const [categories, setCategories] = useState([]);
  const [prompts, setPrompts] = useState([]);
  const [scores, setScores] = useState([]);
  const [calculatedScores, setCalculatedScores] = useState([]);
  const [sortConfig, setSortConfig] = useState({
    key: "",
    direction: "ascending",
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Fetch all required data
        const [modelsData, promptsData, scoresData] = await Promise.all([
          fetch("/api/models").then((res) => res.json()),
          fetch("/api/prompts").then((res) => res.json()),
          fetch("/api/scores").then((res) => res.json()),
        ]);

        // Preprocess models to extract numeric ID from name
        const processedModels = modelsData.map((model) => {
          const match = model.name.match(/Model (\d+)/);
          const derivedId = match ? parseInt(match[1], 10) : null; // Extract number from name
          return { ...model, derivedId };
        });

        // Update states
        setModels(processedModels);
        setPrompts(promptsData);
        setScores(scoresData);

        // Derive unique categories
        const uniqueCategories = [
          ...new Set(
            promptsData.map((prompt) => prompt.category.toLowerCase()),
          ),
        ];
        setCategories(uniqueCategories);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, []); // Run only once on mount

  useEffect(() => {
    if (models.length && scores.length && categories.length && prompts.length) {
      const calculateScores = (modelDerivedId) => {
        const modelScores = scores.filter(
          (score) => Number(score.modelId) === Number(modelDerivedId),
        );

        if (!modelScores.length) return { overall: 0 };

        const categoryScores = {};
        let overallScore = 0;
        let totalPrompts = 0;

        let totalScore = 0;
        let promptCount = 0;
        categories.forEach((category) => {
          const categoryPrompts = prompts.filter((prompt) => {
            const promptScores = modelScores.filter((score) => {
              const match = prompt.content.match(/Prompt (\d+) content/);
              const derivedId = match ? parseInt(match[1], 10) : null; // Extract number from name
              return String(score.promptId) === String(derivedId);
            });
            if (promptScores.length > 0) {
              totalScore += promptScores[0].score;
              promptCount++;
            }
            return prompt.category.toLowerCase() === category;
          });

          categoryScores[category] =
            promptCount > 0 ? (totalScore / promptCount).toFixed(2) : 0;
          overallScore += Number(categoryScores[category]);
          totalPrompts += promptCount;
        });

        overallScore = totalPrompts > 0 ? overallScore.toFixed(2) : 0;
        return { ...categoryScores, overall: overallScore };
      };

      // Calculate scores for all models
      const calculated = models.map((model) => ({
        id: model.id,
        scores: calculateScores(model.derivedId), // Use derived ID
      }));

      setCalculatedScores(calculated);
    }
  }, [models, scores, categories, prompts]); // Dependency array ensures this runs only when all states update

  const sortTable = (key) => {
    console.log(`Sorting by key: ${key}`);
    let direction = "ascending";
    if (sortConfig.key === key && sortConfig.direction === "ascending") {
      direction = "descending";
    }
    setSortConfig({ key, direction });

    const sortedModels = [...calculatedScores].sort((a, b) => {
      if (key === "name") {
        const nameA = models.find(model => model.id === a.id)?.name || '';
        const nameB = models.find(model => model.id === b.id)?.name || '';
        if (nameA < nameB) {
          return direction === "ascending" ? -1 : 1;
        }
        if (nameA > nameB) {
          return direction === "ascending" ? 1 : -1;
        }
        return 0;
      } else {
        const valueA = a.scores[key] || 0;
        const valueB = b.scores[key] || 0;
        if (valueA < valueB) {
          return direction === "ascending" ? -1 : 1;
        }
        if (valueA > valueB) {
          return direction === "ascending" ? 1 : -1;
        }
        return 0;
      }
    });
    setCalculatedScores(sortedModels);
  };

  const handleModelClick = (model) => {
    setSelectedModel(model);
  };

  const handleClosePopup = () => {
    setSelectedModel(null);
  };

  const handleSaveModel = (editedModel) => {
    setModels(
      models.map((model) =>
        model.id === editedModel.id ? editedModel : model,
      ),
    );
  };

  if (!models.length || !categories.length || !prompts.length) {
    return <div>Loading...</div>;
  }

  return (
    <Layout>
      <h1>Leaderboard</h1>
      <table>
        <thead>
          <tr>
            <th onClick={() => sortTable("name")}>Model Name</th>
            {categories.map((category, index) => (
              <th key={index} onClick={() => sortTable(category.toLowerCase())}>
                {category}
              </th>
            ))}
            <th onClick={() => sortTable("overall")}>Overall Score</th>
          </tr>
        </thead>
        <tbody>
          {calculatedScores.map(({ id, scores }) => {
            const model = models.find((m) => m.id === id);
            if (!model) return null; // Skip if model not found
            return (
              <tr key={id} onClick={() => handleModelClick(model)}>
                <td>{model.name}</td>
                {categories.map((category, idx) => (
                  <td key={idx}>{scores[category] || 0}</td>
                ))}
                <td>{scores.overall || 0}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
      {selectedModel && (
        <InputPopup
          item={selectedModel}
          onClose={handleClosePopup}
          onSave={handleSaveModel}
          categories={categories}
          prompts={prompts}
        />
      )}
    </Layout>
  );
}
