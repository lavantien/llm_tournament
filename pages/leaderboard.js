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
    // Fetch models from the API route
    fetch('/api/mocks/models')
      .then(response => response.json())
      .then(data => setModels(data))
      .catch(error => console.error('Error fetching models:', error));

    // Fetch prompts from the API route
    fetch('/api/mocks/prompts')
      .then(response => response.json())
      .then(data => {
        setPrompts(data);
        const uniqueCategories = [...new Set(data.map(prompt => prompt.category))];
        setCategories(uniqueCategories);
      })
      .catch(error => console.error('Error fetching prompts:', error));

    // Fetch scores from the API route
    fetch('/api/mocks/scores')
      .then(response => response.json())
      .then(data => setScores(data))
      .catch(error => console.error('Error fetching scores:', error));
  }, []);

  const calculateScores = (modelId) => {
    const modelScores = scores.filter(score => score.modelId === modelId);
    const categoryScores = {};
    let overallScore = 0;

    categories.forEach(category => {
      const categoryPrompts = prompts.filter(prompt => prompt.category === category);
      let totalScore = 0;
      let promptCount = 0;

      categoryPrompts.forEach(prompt => {
        const promptScores = modelScores.filter(score => score.promptId === prompt.id);
        if (promptScores.length > 0) {
          totalScore += promptScores[0].score;
          promptCount++;
        }
      });

      if (promptCount > 0) {
        categoryScores[category.toLowerCase()] = (totalScore / promptCount).toFixed(2);
        overallScore += totalScore / promptCount;
      }
    });

    overallScore = (overallScore / categories.length).toFixed(2);
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
    setModels(models.map(model => (model.name === editedModel.name ? editedModel : model)));
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
          {models.map((model, index) => (
            <tr key={index} onClick={() => handleModelClick(model)}>
              <td>{model.name}</td>
              {categories.map((category, idx) => (
                <td key={idx}>{calculateScores(model.id)[category.toLowerCase()]}</td>
              ))}
              <td>{calculateScores(model.id).overall}</td>
            </tr>
          ))}
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
