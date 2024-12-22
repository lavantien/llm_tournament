import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';

export default function ModelManager() {
  const [models, setModels] = useState([]);
  const [formData, setFormData] = useState({
    name: '',
    path: '',
    gpuLayers: '',
    ctxSize: '',
    batchSize: '',
    threads: '',
    keep: '',
    predict: '',
    flashAttn: false,
    mlock: false,
    cacheTypeK: '',
    cacheTypeV: ''
  });
  const [selectedModel, setSelectedModel] = useState(null);

  useEffect(() => {
    // Fetch models from the JSON file
    fetch('/mocks/models.json')
      .then(response => response.json())
      .then(data => setModels(data))
      .catch(error => console.error('Error fetching models:', error));
  }, []);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? checked : value
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission (e.g., add model to the database)
    setModels([...models, formData]);
    setFormData({
      name: '',
      path: '',
      gpuLayers: '',
      ctxSize: '',
      batchSize: '',
      threads: '',
      keep: '',
      predict: '',
      flashAttn: false,
      mlock: false,
      cacheTypeK: '',
      cacheTypeV: ''
    });
  };

  const handleModelClick = (model) => {
    setSelectedModel(model);
  };

  const handleClosePopup = () => {
    setSelectedModel(null);
  };

  return (
    <Layout>
      <h1>Model Manager</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name:</label>
          <input type="text" name="name" value={formData.name} onChange={handleChange} />
        </div>
        <div>
          <label>Path:</label>
          <input type="text" name="path" value={formData.path} onChange={handleChange} />
        </div>
        <div>
          <label>GPU Layers:</label>
          <input type="number" name="gpuLayers" value={formData.gpuLayers} onChange={handleChange} />
        </div>
        <div>
          <label>Context Size:</label>
          <input type="number" name="ctxSize" value={formData.ctxSize} onChange={handleChange} />
        </div>
        <div>
          <label>Batch Size:</label>
          <input type="number" name="batchSize" value={formData.batchSize} onChange={handleChange} />
        </div>
        <div>
          <label>Threads:</label>
          <input type="number" name="threads" value={formData.threads} onChange={handleChange} />
        </div>
        <div>
          <label>Keep:</label>
          <input type="number" name="keep" value={formData.keep} onChange={handleChange} />
        </div>
        <div>
          <label>Predict:</label>
          <input type="number" name="predict" value={formData.predict} onChange={handleChange} />
        </div>
        <div>
          <label>Flash Attn:</label>
          <input type="checkbox" name="flashAttn" checked={formData.flashAttn} onChange={handleChange} />
        </div>
        <div>
          <label>Mlock:</label>
          <input type="checkbox" name="mlock" checked={formData.mlock} onChange={handleChange} />
        </div>
        <div>
          <label>Cache Type K:</label>
          <input type="text" name="cacheTypeK" value={formData.cacheTypeK} onChange={handleChange} />
        </div>
        <div>
          <label>Cache Type V:</label>
          <input type="text" name="cacheTypeV" value={formData.cacheTypeV} onChange={handleChange} />
        </div>
        <button type="submit">Add Model</button>
      </form>
      <ul>
        {models.map((model, index) => (
          <li key={index} onClick={() => handleModelClick(model)}>
            {model.name}
          </li>
        ))}
      </ul>
      <DetailPopup item={selectedModel} onClose={handleClosePopup} />
    </Layout>
  );
}
