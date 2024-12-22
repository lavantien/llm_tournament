import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import InputPopup from '../components/InputPopup';

export default function Leaderboard() {
  const [models, setModels] = useState([]);
  const [selectedModel, setSelectedModel] = useState(null);
  const [categories, setCategories] = useState([]);
  const [prompts, setPrompts] = useState([]);
  const [scores, setScores] = useState([]);

  useEffect(() => {
    // Fetch models from the database
    fetch('/api/models')
      .then(response => response.json())
      .then(data => {
        console.log('Fetched models:', data);
        setModels(data);
      })
      .catch(error => console.error('Error fetching models:', error));

    // Fetch prompts from the database
    fetch('/api/prompts')
      .then(response => response.json())
      .then(data => {
        console.log('Fetched prompts:', data);
        setPrompts(data);
        const uniqueCategories = [...new Set(data.map(prompt => prompt.category))];
        setCategories(uniqueCategories);
      })
      .catch(error => console.error('Error fetching prompts:', error));

    // Fetch scores from the database
    fetch('/api/scores')
      .then(response => response.json())
      .then(data => {
        console.log('Fetched scores:', data);
        setScores(data);
      })
      .catch(error => console.error('Error fetching scores:', error));
  }, []);

  const calculateScores = (modelId) => {
    console.log('Calculating scores for modelId:', modelId);
    const modelScores = scores.filter(score => score.modelId === modelId);
    console.log('Model scores for modelId', modelId, ':', modelScores);
    const categoryScores = {};
    let overallScore = 0;

    categories.forEach(category => {
      const categoryPrompts = prompts.filter(prompt => prompt.category === category);
      console.log('Category prompts for', category, ':', categoryPrompts);
      let totalScore = 0;
      let promptCount = 0;

      categoryPrompts.forEach(prompt => {
        const promptScores = modelScores.filter(score => score.promptId === prompt.id);
        console.log('Prompt scores for promptId', prompt.id, ':', promptScores);
        if (promptScores.length > 0) {
          totalScore += promptScores[0].score;
          promptCount++;
        }
      });

      categoryScores[category.toLowerCase()] = promptCount > 0 ? (totalScore / promptCount).toFixed(2) : 0;
      console.log('Category score for', category, ':', categoryScores[category.toLowerCase()]);
      overallScore += totalScore;
    });

    overallScore = (overallScore / (categories.length * prompts.length)).toFixed(2);
    console.log('Overall score for modelId', modelId, ':', overallScore);
    console.log('Calculated scores for modelId', modelId, ':', { ...categoryScores, overall: overallScore });
    return { ...categoryScores, overall: overallScore };
  };

  const sortTable = (key) => {
    const sortedModels = [...models].sort((a, b) => (a[key] < b[key] ? 1 : -1));
    setModels(sortedModels);
  };

  const handleModelClick = (model) => {
    setSelectedModel(model);
  };

  const handleClosePopup = () => {
    setSelectedModel(null);
  };

  const handleSaveModel = (editedModel) => {
    setModels(models.map(model => (model.id === editedModel.id ? editedModel : model)));
  };

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
          {models.map((model, index) => {
            const scores = calculateScores(model.id);
            return (
              <tr key={index} onClick={() => handleModelClick(model)}>
                <td>{model.name}</td>
                {categories.map((category, idx) => (
                  <td key={idx}>{scores[category.toLowerCase()]}</td>
                ))}
                <td>{scores.overall}</td>
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
