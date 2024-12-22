import { useState, useEffect, useMemo } from "react";
import Layout from "../components/Layout";
import InputPopup from "../components/InputPopup";
import { Table, Container } from 'react-bootstrap';

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

  const memoizedScores = useMemo(() => {
    if (models.length && scores.length && categories.length && prompts.length) {
      const calculateScores = (modelDerivedId) => {
        const modelScores = scores.filter(
          (score) => Number(score.modelId) === Number(modelDerivedId),
        );

        console.log(
          `ModelScores for modelDerivedId ${modelDerivedId}:`,
          modelScores,
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

      console.log("Calculated scores:", calculated);
      return calculated;
    }
    return [];
  }, [models, scores, categories, prompts]); // Dependency array ensures this runs only when all states update

  useEffect(() => {
    setCalculatedScores(memoizedScores);
  }, [memoizedScores]);

  const sortTable = (key) => {
    console.log(`Sorting by key: ${key}`);
    let direction = "ascending";
    if (sortConfig.key === key && sortConfig.direction === "ascending") {
      direction = "descending";
    }
    setSortConfig({ key, direction });

    const sortedModels = [...calculatedScores].sort((a, b) => {
      if (key === "name") {
        const nameA = models.find((model) => model.id === a.id)?.name || "";
        const nameB = models.find((model) => model.id === b.id)?.name || "";
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
      <Container>
        <h1 className="h3 mb-4">Leaderboard</h1>
        <Table striped bordered hover variant="dark">
          <thead>
            <tr>
              <th className="cursor-pointer" onClick={() => sortTable("name")}>Model Name</th>
              {categories.map((category, index) => (
                <th key={index} className="cursor-pointer" onClick={() => sortTable(category.toLowerCase())}>
                  {category}
                </th>
              ))}
              <th className="cursor-pointer" onClick={() => sortTable("overall")}>Overall Score</th>
            </tr>
          </thead>
          <tbody>
            {calculatedScores.map(({ id, scores }) => {
              const model = models.find((m) => m.id === id);
              if (!model) return null; // Skip if model not found
              return (
                <tr key={id} className="cursor-pointer" onClick={() => handleModelClick(model)}>
                  <td>{model.name}</td>
                  {categories.map((category, idx) => (
                    <td key={idx}>{scores[category] || 0}</td>
                  ))}
                  <td>{scores.overall || 0}</td>
                </tr>
              );
            })}
          </tbody>
        </Table>
        {selectedModel && (
          <InputPopup
            item={selectedModel}
            onClose={handleClosePopup}
            onSave={handleSaveModel}
            categories={categories}
            prompts={prompts}
            scores={scores}
          />
        )}
      </Container>
    </Layout>
  );
}
