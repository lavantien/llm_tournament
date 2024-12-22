import { useState, useEffect } from 'react';
import Layout from '../components/Layout';

export default function Leaderboard() {
  const [models, setModels] = useState([
    { name: 'Model 1', category1: 90, category2: 85, overall: 87 },
    { name: 'Model 2', category1: 88, category2: 92, overall: 90 },
    { name: 'Model 3', category1: 95, category2: 80, overall: 87 },
    { name: 'Model 4', category1: 78, category2: 89, overall: 83 },
    { name: 'Model 5', category1: 92, category2: 77, overall: 84 }
    // Add more mock data as needed
  ]);

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
  }, []);

  const sortTable = (key) => {
    const sortedModels = [...models].sort((a, b) => (a[key] < b[key] ? 1 : -1));
    setModels(sortedModels);
  };

  return (
    <Layout>
      <h1>Leaderboard</h1>
      <table>
        <thead>
          <tr>
            <th onClick={() => sortTable('name')}>Model Name</th>
            <th onClick={() => sortTable('category1')}>Category 1</th>
            <th onClick={() => sortTable('category2')}>Category 2</th>
            <th onClick={() => sortTable('overall')}>Overall Score</th>
          </tr>
        </thead>
        <tbody>
          {models.map((model, index) => (
            <tr key={index} onClick={() => alert(`Details for ${model.name}`)}>
              <td>{model.name}</td>
              <td>{model.category1}</td>
              <td>{model.category2}</td>
              <td>{model.overall}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </Layout>
  );
}
