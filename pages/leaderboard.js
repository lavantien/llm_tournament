import { useState, useEffect, useMemo } from 'react';
import Layout from '../components/Layout';
import InputPopup from '../components/InputPopup';

export default function Leaderboard() {
  const [models, setModels] = useState([]);
  const [selectedModel, setSelectedModel] = useState(null);
  const [categories, setCategories] = useState([]);
  const [prompts, setPrompts] = useState([]);
  const [scores, setScores] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const fetchModels = fetch('/api/models')
          .then(response => {
            if (!response.ok) throw new Error('Failed to fetch models');
            return response.json();
          });

        const fetchPrompts = fetch('/api/prompts')
          .then(response => {
            if (!response.ok) throw new Error('Failed to fetch prompts');
            return response.json();
          });

        const fetchScores = fetch('/api/scores')
          .then(response => {
            if (!response.ok) throw new Error('Failed to fetch scores');
            return response.json();
          });

        const [modelsData, promptsData, scoresData] = await Promise.all([fetchModels, fetchPrompts, fetchScores]);

        console.log('Fetched models:', modelsData);
        console.log('Fetched prompts:', promptsData);
        console.log('Fetched scores:', scoresData);

        setModels(modelsData);
        console.log('Models set:', modelsData);
        setPrompts(promptsData);
        console.log('Prompts set:', promptsData);
        setScores(scoresData);
        console.log('Scores set:', scoresData);

        const uniqueCategories = [...new Set(promptsData.map(prompt => prompt.category.toLowerCase()))];
        setCategories(uniqueCategories);
        console.log('Categories set:', uniqueCategories);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);

  const memoizedScores = useMemo(() => {
    if (!models.length || !scores.length || !categories.length || !prompts.length) {
      console.log('Memoized scores skipped due to missing data');
      return [];
    }

    const calculateScores = (modelId) => {
      console.log('Calculating scores for modelId:', modelId);
      const modelScores = scores.filter(score => score.modelId === modelId);
      console.log('Model scores for modelId', modelId, ':', modelScores);

      if (modelScores.length === 0) return { overall: 0 };

      const categoryScores = {};
      let overallScore = 0;
      let totalPrompts = 0;

      categories.forEach(category => {
        const categoryPrompts = prompts.filter(prompt => prompt.category.toLowerCase() === category);
        console.log('Category prompts for', category, ':', categoryPrompts);
        let totalScore = 0;
        let promptCount = 0;

        categoryPrompts.forEach(prompt => {
          const promptScores = modelScores.filter(score => `${score.promptId}` === `${prompt.id}`);
          console.log('Prompt scores for promptId', prompt.id, ':', promptScores);
          if (promptScores.length > 0) {
            totalScore += promptScores[0].score;
            promptCount++;
          }
        });

        categoryScores[category] = promptCount > 0 ? (totalScore / promptCount).toFixed(2) : 0;
        console.log('Category score for', category, ':', categoryScores[category]);
        overallScore += totalScore;
        totalPrompts += promptCount;
      });

      overallScore = totalPrompts > 0 ? (overallScore / totalPrompts).toFixed(2) : 0;
      console.log('Overall score for modelId', modelId, ':', overallScore);
      console.log('Calculated scores for modelId', modelId, ':', { ...categoryScores, overall: overallScore });
      return { ...categoryScores, overall: overallScore };
    };

    return models.map(model => ({
      id: model.id,
      scores: calculateScores(model.id),
    }));
  }, [models, scores, categories, prompts]);

  console.log('Memoized scores:', memoizedScores);

  const sortTable = (key) => {
    const sortedModels = [...models].sort((a, b) => ((a[key] || 0) < (b[key] || 0) ? 1 : -1));
    setModels(sortedModels);
    console.log('Models sorted by', key, ':', sortedModels);
  };

  const handleModelClick = (model) => {
    setSelectedModel(model);
    console.log('Selected model set:', model);
  };

  const handleClosePopup = () => {
    setSelectedModel(null);
    console.log('Selected model reset');
  };

  const handleSaveModel = (editedModel) => {
    setModels(models.map(model => (model.id === editedModel.id ? editedModel : model)));
    console.log('Models set:', models.map(model => (model.id === editedModel.id ? editedModel : model)));
  };

  if (!models.length || !scores.length || !categories.length || !prompts.length) {
    return <div>Loading...</div>;
  }

  return (
    <Layout>
      <h1>Leaderboard</h1>
      <table>
        <thead>
          <tr>
            <th onClick={() => sortTable('name')}>Model Name</th>
            {categories.map((category, index) => (
              <th key={index} onClick={() => sortTable(category.toLowerCase())}>{category}</th>
            ))}
            <th onClick={() => sortTable('overall')}>Overall Score</th>
          </tr>
        </thead>
        <tbody>
          {memoizedScores.map(({ id, scores }) => {
            const model = models.find(m => m.id === id);
            if (!model) return null; // Skip if model not found
            console.log('Rendering model:', model);
            console.log('Rendering scores:', scores);
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
          scores={scores}
        />
      )}
    </Layout>
  );
}
