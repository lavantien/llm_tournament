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
    // Fetch models from the API route
    fetch('/api/models')
      .then(response => response.json())
      .then(data => {
        setModels(data);
        console.log('Models set:', data);
      })
      .catch(error => console.error('Error fetching models:', error));

    // Listen for dataWiped event
    const handleDataWiped = () => {
      setModels([]);
      console.log('Data wiped, models reset');
    };

    window.addEventListener('dataWiped', handleDataWiped);

    return () => {
      window.removeEventListener('dataWiped', handleDataWiped);
    };
  }, []);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? checked : value
    });
    console.log('Form data set:', { ...formData, [name]: type === 'checkbox' ? checked : value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission (e.g., add model to the database)
    setModels([...models, formData]);
    console.log('Models set:', [...models, formData]);
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
    console.log('Form data reset');
  };

  const handleModelClick = (model) => {
    setSelectedModel(model);
    console.log('Selected model set:', model);
  };

  const handleClosePopup = () => {
    setSelectedModel(null);
    console.log('Selected model reset');
  };

  return (
    <Layout>
      <h1 className="text-3xl font-bold mb-4">Model Manager</h1>
      <form onSubmit={handleSubmit} className="bg-mystic-secondary p-4 rounded shadow-lg">
        <div className="mb-4">
          <label className="font-semibold">Name:</label>
          <input type="text" name="name" value={formData.name} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Path:</label>
          <input type="text" name="path" value={formData.path} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">GPU Layers:</label>
          <input type="number" name="gpuLayers" value={formData.gpuLayers} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Context Size:</label>
          <input type="number" name="ctxSize" value={formData.ctxSize} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Batch Size:</label>
          <input type="number" name="batchSize" value={formData.batchSize} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Threads:</label>
          <input type="number" name="threads" value={formData.threads} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Keep:</label>
          <input type="number" name="keep" value={formData.keep} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Predict:</label>
          <input type="number" name="predict" value={formData.predict} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Flash Attn:</label>
          <input type="checkbox" name="flashAttn" checked={formData.flashAttn} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Mlock:</label>
          <input type="checkbox" name="mlock" checked={formData.mlock} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Cache Type K:</label>
          <input type="text" name="cacheTypeK" value={formData.cacheTypeK} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Cache Type V:</label>
          <input type="text" name="cacheTypeV" value={formData.cacheTypeV} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <button type="submit" className="bg-mystic-highlight text-mystic-primary px-4 py-2 rounded hover:bg-mystic-accent">Add Model</button>
      </form>
      <ul className="list-disc pl-4">
        {models.map((model, index) => (
          <li key={index} className="cursor-pointer hover:text-mystic-highlight" onClick={() => handleModelClick(model)}>
            {model.name}
          </li>
        ))}
      </ul>
      <DetailPopup item={selectedModel} onClose={handleClosePopup} />
    </Layout>
  );
}
