import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';
import { Form, Button, Container, ListGroup } from 'react-bootstrap';

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
      <Container>
        <h1 className="h3 mb-4">Model Manager</h1>
        <Form onSubmit={handleSubmit} className="bg-secondary p-3 rounded shadow">
          <Form.Group controlId="formName">
            <Form.Label>Name:</Form.Label>
            <Form.Control type="text" name="name" value={formData.name} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formPath">
            <Form.Label>Path:</Form.Label>
            <Form.Control type="text" name="path" value={formData.path} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formGpuLayers">
            <Form.Label>GPU Layers:</Form.Label>
            <Form.Control type="number" name="gpuLayers" value={formData.gpuLayers} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formCtxSize">
            <Form.Label>Context Size:</Form.Label>
            <Form.Control type="number" name="ctxSize" value={formData.ctxSize} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formBatchSize">
            <Form.Label>Batch Size:</Form.Label>
            <Form.Control type="number" name="batchSize" value={formData.batchSize} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formThreads">
            <Form.Label>Threads:</Form.Label>
            <Form.Control type="number" name="threads" value={formData.threads} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formKeep">
            <Form.Label>Keep:</Form.Label>
            <Form.Control type="number" name="keep" value={formData.keep} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formPredict">
            <Form.Label>Predict:</Form.Label>
            <Form.Control type="number" name="predict" value={formData.predict} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formFlashAttn">
            <Form.Check type="checkbox" label="Flash Attn:" name="flashAttn" checked={formData.flashAttn} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formMlock">
            <Form.Check type="checkbox" label="Mlock:" name="mlock" checked={formData.mlock} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formCacheTypeK">
            <Form.Label>Cache Type K:</Form.Label>
            <Form.Control type="text" name="cacheTypeK" value={formData.cacheTypeK} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formCacheTypeV">
            <Form.Label>Cache Type V:</Form.Label>
            <Form.Control type="text" name="cacheTypeV" value={formData.cacheTypeV} onChange={handleChange} />
          </Form.Group>
          <Button type="submit">Add Model</Button>
        </Form>
        <ListGroup>
          {models.map((model, index) => (
            <ListGroup.Item key={index} action onClick={() => handleModelClick(model)}>
              {model.name}
            </ListGroup.Item>
          ))}
        </ListGroup>
        <DetailPopup item={selectedModel} onClose={handleClosePopup} />
      </Container>
    </Layout>
  );
}
