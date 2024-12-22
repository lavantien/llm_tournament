import React, { useState, useEffect } from 'react';
import styles from '../styles/DetailPopup.module.css';

const InputPopup = ({ item, onClose, onSave, categories, prompts }) => {
  const [editedItem, setEditedItem] = useState({ ...item });
  const [selectedCategory, setSelectedCategory] = useState('');
  const [selectedPrompt, setSelectedPrompt] = useState('');
  const [score, setScore] = useState('');

  useEffect(() => {
    setEditedItem({ ...item });
  }, [item]);

  const handleCategoryChange = (e) => {
    setSelectedCategory(e.target.value);
    setSelectedPrompt('');
    setScore('');
  };

  const handlePromptChange = (e) => {
    setSelectedPrompt(e.target.value);
    setScore('');
  };

  const handleScoreChange = (e) => {
    setScore(e.target.value);
  };

  const handleSave = () => {
    const categoryKey = selectedCategory.toLowerCase();
    const newScore = parseFloat(score);
    const currentScore = editedItem[categoryKey] || 0;
    const newOverallScore =
      ((Object.keys(editedItem)
        .filter((key) => key.startsWith('category'))
        .reduce((sum, key) => sum + (editedItem[key] || 0), 0) +
        newScore -
        currentScore) /
        (categories.length + 1)) *
      100;

    const updatedItem = {
      ...editedItem,
      [categoryKey]: newScore,
      overall: newOverallScore.toFixed(2)
    };

    onSave(updatedItem);
    onClose();
  };

  const filteredPrompts = prompts.filter(prompt => prompt.category === selectedCategory);

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
            <label>Score:</label>
            <input type="number" value={score} onChange={handleScoreChange} />
          </div>
        )}
        <button onClick={handleSave}>Save</button>
      </div>
    </div>
  );
};

export default InputPopup;
