import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Table } from 'react-bootstrap';

const InputPopup = ({ item, onClose, onSave, categories, prompts = [], scores = [] }) => {
  const [editedItem, setEditedItem] = useState({ ...item });
  const [selectedCategory, setSelectedCategory] = useState('');
  const [selectedPrompt, setSelectedPrompt] = useState('');
  const [attempts, setAttempts] = useState(1);

  useEffect(() => {
    setEditedItem({ ...item });
    console.log('Edited item set:', { ...item });
  }, [item]);

  const handleCategoryChange = (e) => {
    setSelectedCategory(e.target.value);
    setSelectedPrompt('');
    setAttempts(1);
    console.log('Selected category set:', e.target.value);
  };

  const handlePromptChange = (e) => {
    setSelectedPrompt(e.target.value);
    setAttempts(1);
    console.log('Selected prompt set:', e.target.value);
  };

  const handleAttemptsChange = (e) => {
    setAttempts(parseInt(e.target.value));
    console.log('Attempts set:', parseInt(e.target.value));
  };

  const handleSave = async () => {
    const categoryKey = selectedCategory.toLowerCase();
    const prompt = prompts.find(p => p.content === selectedPrompt);
    const modelId = item.id;
    const promptId = prompt.id;
    let score = 0;

    if (attempts === 1) {
      score = 100;
    } else if (attempts === 2) {
      score = 50;
    } else if (attempts === 3) {
      score = 20;
    }

    const response = await fetch('/api/scores', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ modelId, promptId, attempts, score })
    });

    if (response.ok) {
      const updatedItem = {
        ...editedItem,
        [categoryKey]: score,
        overall: calculateOverallScore(editedItem, categoryKey, score)
      };

      onSave(updatedItem);
      console.log('Item saved:', updatedItem);
      onClose();
    } else {
      console.error('Error saving score:', await response.json());
    }
  };

  const calculateOverallScore = (item, categoryKey, newScore) => {
    const currentScore = item[categoryKey] || 0;
    const newOverallScore =
      ((Object.keys(item)
        .filter((key) => key.startsWith('category'))
        .reduce((sum, key) => sum + (item[key] || 0), 0) +
        newScore -
        currentScore) /
        (categories.length + 1)) *
      100;
    return newOverallScore.toFixed(2);
  };

  const filteredPrompts = selectedCategory ? prompts.filter(prompt => prompt.category === selectedCategory) : [];
  const modelScores = scores.filter(score => score.modelId === item.id);

  return (
    <Modal show onHide={onClose} centered>
      <Modal.Header closeButton>
        <Modal.Title>Edit Scores</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form.Group controlId="formCategory">
          <Form.Label>Category:</Form.Label>
          <Form.Select value={selectedCategory} onChange={handleCategoryChange}>
            <option value="">Select a category</option>
            {categories.map((category, index) => (
              <option key={index} value={category}>
                {category}
              </option>
            ))}
          </Form.Select>
        </Form.Group>
        {selectedCategory && (
          <Form.Group controlId="formPrompt">
            <Form.Label>Prompt:</Form.Label>
            <Form.Select value={selectedPrompt} onChange={handlePromptChange}>
              <option value="">Select a prompt</option>
              {filteredPrompts.map((prompt, index) => (
                <option key={index} value={prompt.content}>
                  {prompt.content}
                </option>
              ))}
            </Form.Select>
          </Form.Group>
        )}
        {selectedPrompt && (
          <Form.Group controlId="formAttempts">
            <Form.Label>Attempts:</Form.Label>
            <Form.Control type="number" value={attempts} onChange={handleAttemptsChange} min="1" />
          </Form.Group>
        )}
        <Button onClick={handleSave}>Save</Button>
        <h3 className="h4 mt-4 mb-2">Scores for {item.name}</h3>
        <Table striped bordered hover>
          <thead>
            <tr>
              <th>Prompt</th>
              <th>Score</th>
            </tr>
          </thead>
          <tbody>
            {modelScores.map((score, index) => {
              const prompt = prompts.find(p => p.id === score.promptId);
              return (
                <tr key={index}>
                  <td>{prompt ? prompt.content : 'Unknown Prompt'}</td>
                  <td>{score.score}</td>
                </tr>
              );
            })}
          </tbody>
        </Table>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={onClose}>
          Close
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default InputPopup;
