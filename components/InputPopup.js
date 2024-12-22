import React, { useState, useEffect } from 'react';
import styles from '../styles/DetailPopup.module.css';

const InputPopup = ({ item, onClose, onSave, categories, prompts, scores }) => {
  const [editedItem, setEditedItem] = useState({ ...item });
  const [selectedCategory, setSelectedCategory] = useState('');
  const [selectedPrompt, setSelectedPrompt] = useState('');
  const [attempts, setAttempts] = useState(1);

  useEffect(() => {
    setEditedItem({ ...item });
  }, [item]);

  const handleCategoryChange = (e) => {
    setSelectedCategory(e.target.value);
    setSelectedPrompt('');
    setAttempts(1);
  };

  const handlePromptChange = (e) => {
    setSelectedPrompt(e.target.value);
    setAttempts(1);
  };

  const handleAttemptsChange = (e) => {
    setAttempts(parseInt(e.target.value));
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

  const filteredPrompts = prompts.filter(prompt => prompt.category === selectedCategory);

  const modelScores = scores.filter(score => score.modelId === item.id);

  return (
    <div className={styles.popupOverlay}>
      <div className={styles.popupContent}>
        <button className={styles.closeButton} onClick={onClose}>
          &times;
        </button>
        <h2>Edit Scores</h2>
        <div>
          <label>Category:</label>
          <select value={selectedCategory} onChange={handleCategoryChange}>
            <option value="">Select a category</option>
            {categories.map((category, index) => (
              <option key={index} value={category}>
                {category}
              </option>
            ))}
          </select>
        </div>
        {selectedCategory && (
          <div>
            <label>Prompt:</label>
            <select value={selectedPrompt} onChange={handlePromptChange}>
              <option value="">Select a prompt</option>
              {filteredPrompts.map((prompt, index) => (
                <option key={index} value={prompt.content}>
                  {prompt.content}
                </option>
              ))}
            </select>
          </div>
        )}
        {selectedPrompt && (
          <div>
            <label>Attempts:</label>
            <input type="number" value={attempts} onChange={handleAttemptsChange} min="1" />
          </div>
        )}
        <button onClick={handleSave}>Save</button>
        <h3>Scores for {item.name}</h3>
        <table>
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
        </table>
      </div>
    </div>
  );
};

export default InputPopup;
