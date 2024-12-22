import React, { useState, useEffect } from 'react';
import styles from '../styles/DetailPopup.module.css';

const DetailPopup = ({ item, onClose, onSave, categories, prompts }) => {
  const [editedItem, setEditedItem] = useState({ ...item });
  const [selectedCategory, setSelectedCategory] = useState('');
  const [selectedPrompt, setSelectedPrompt] = useState('');
  const [score, setScore] = useState('');

  useEffect(() => {
    if (selectedCategory) {
      setSelectedPrompt('');
      setScore('');
    }
  }, [selectedCategory]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setEditedItem({
      ...editedItem,
      [name]: value
    });
  };

  const handleCategoryChange = (e) => {
    setSelectedCategory(e.target.value);
  };

  const handlePromptChange = (e) => {
    setSelectedPrompt(e.target.value);
  };

  const handleScoreChange = (e) => {
    setScore(e.target.value);
  };

  const handleSave = () => {
    if (selectedCategory && selectedPrompt && score) {
      const categoryKey = selectedCategory.toLowerCase();
      const currentScore = editedItem[categoryKey] || 0;
      const newScore = currentScore + parseInt(score, 10);
      const overallScore = Object.keys(editedItem)
        .filter(key => key.startsWith('category'))
        .reduce((sum, key) => sum + (editedItem[key] || 0), 0) + parseInt(score, 10);

      setEditedItem({
        ...editedItem,
        [categoryKey]: newScore,
        overall: overallScore
      });

      onSave({
        ...editedItem,
        [categoryKey]: newScore,
        overall: overallScore
      });
    }
    onClose();
  };

  if (!item) return null;

  const filteredPrompts = prompts.filter(prompt => prompt.category === selectedCategory);

  return (
    <div className={styles.popupOverlay}>
      <div className={styles.popupContent}>
        <button className={styles.closeButton} onClick={onClose}>Close</button>
        <h2>Details</h2>
        <div>
          <label>Category:</label>
          <select name="category" value={selectedCategory} onChange={handleCategoryChange}>
            <option value="">Select a category</option>
            {categories.map((category, index) => (
              <option key={index} value={category}>{category}</option>
            ))}
          </select>
        </div>
        {selectedCategory && (
          <div>
            <label>Prompt:</label>
            <select name="prompt" value={selectedPrompt} onChange={handlePromptChange}>
              <option value="">Select a prompt</option>
              {filteredPrompts.map((prompt, index) => (
                <option key={index} value={prompt.content}>{prompt.content}</option>
              ))}
            </select>
          </div>
        )}
        {selectedPrompt && (
          <div>
            <label>Score:</label>
            <input type="number" name="score" value={score} onChange={handleScoreChange} />
          </div>
        )}
        {Object.keys(editedItem).filter(key => key.startsWith('category')).map((key, index) => (
          <div key={index}>
            <label>{key}:</label>
            <input type="number" name={key} value={editedItem[key]} onChange={handleChange} readOnly />
          </div>
        ))}
        <button onClick={handleSave}>Save</button>
      </div>
    </div>
  );
};

export default DetailPopup;
