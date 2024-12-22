import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';

export default function Leaderboard() {
  const [models, setModels] = useState([
    { name: 'Model 1', category1: 90, category2: 85, overall: 87 },
    { name: 'Model 2', category1: 88, category2: 92, overall: 90 },
    { name: 'Model 3', category1: 95, category2: 80, overall: 87 },
    { name: 'Model 4', category1: 78, category2: 89, overall: 83 },
    { name: 'Model 5', category1: 92, category2: 77, overall: 84 }
    // Add more mock data as needed
  ]);
  const [selectedModel, setSelectedModel] = useState(null);
  const [categories, setCategories] = useState([]);
  const [prompts, setPrompts] = useState([]);

  useEffect(() => {
    // Fetch models from the database
    // For now, we'll use dummy data
    setModels([
      { name: 'Model 1', category1: 90, category2: 85, overall: 87 },
      { name: 'Model 2', category1: 88, category2: 92, overall: 90 },
      { name: 'Model 3', category1: 95, category2: 80, overall: 87 },
      { name: 'Model 4', category1: 78, category2: 89, overall: 83 },
      { name: 'Model 5', category1: 92, category2: 77, overall: 84 }
    ]);

    // Fetch categories from prompts
    const promptsData = [
      { content: 'Prompt 1 content', solution: 'Prompt 1 solution', profile: 'Profile 1', category: 'Category 1' },
      { content: 'Prompt 2 content', solution: 'Prompt 2 solution', profile: 'Profile 2', category: 'Category 2' },
      { content: 'Prompt 3 content', solution: 'Prompt 3 solution', profile: 'Profile 3', category: 'Category 3' }
      // Add more mock data as needed
    ];
    const uniqueCategories = [...new Set(promptsData.map(prompt => prompt.category))];
    setCategories(uniqueCategories);
    setPrompts(promptsData);
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
      <DetailPopup
        item={selectedModel}
        onClose={handleClosePopup}
        onSave={handleSaveModel}
        categories={categories}
        prompts={prompts}
      />
    </Layout>
  );
}
