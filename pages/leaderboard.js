import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import InputPopup from '../components/InputPopup';

export default function Leaderboard() {
  const [models, setModels] = useState([]);
  const [selectedModel, setSelectedModel] = useState(null);
  const [categories, setCategories] = useState([]);
  const [prompts, setPrompts] = useState([]);

  useEffect(() => {
    // Fetch models from the JSON file
    fetch('/mocks/models.json')
      .then(response => response.json())
      .then(data => setModels(data))
      .catch(error => console.error('Error fetching models:', error));

    // Fetch prompts from the JSON file
    fetch('/mocks/prompts.json')
      .then(response => response.json())
      .then(data => {
        setPrompts(data);
        const uniqueCategories = [...new Set(data.map(prompt => prompt.category))];
        setCategories(uniqueCategories);
      })
      .catch(error => console.error('Error fetching prompts:', error));
  }, []);

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
                <td key={idx}>{model[category.toLowerCase()]}</td>
              ))}
              <td>{model.overall}</td>
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
        />
      )}
    </Layout>
  );
}
