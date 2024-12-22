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

  const filteredPrompts = prompts ? prompts.filter(prompt => prompt.category === selectedCategory) : [];

  return (
    <div className={styles.popupOverlay}>
      <div className={styles.popupContent}>
        <button className={styles.closeButton} onClick={onClose}>Close</button>
        <h2>Details</h2>
        {item.name && (
          <div>
            <label>Name:</label>
            <input type="text" name="name" value={editedItem.name} onChange={handleChange} readOnly />
          </div>
        )}
        {item.content && (
          <div>
            <label>Content:</label>
            <textarea name="content" value={editedItem.content} onChange={handleChange} readOnly />
          </div>
        )}
        {item.solution && (
          <div>
            <label>Solution:</label>
            <textarea name="solution" value={editedItem.solution} onChange={handleChange} readOnly />
          </div>
        )}
        {item.profile && (
          <div>
            <label>Profile:</label>
            <input type="text" name="profile" value={editedItem.profile} onChange={handleChange} readOnly />
          </div>
        )}
        {item.category && (
          <div>
            <label>Category:</label>
            <input type="text" name="category" value={editedItem.category} onChange={handleChange} readOnly />
          </div>
        )}
        {item.systemPrompt && (
          <div>
            <label>System Prompt:</label>
            <textarea name="systemPrompt" value={editedItem.systemPrompt} onChange={handleChange} readOnly />
          </div>
        )}
        {item.dryMultiplier && (
          <div>
            <label>Dry Multiplier:</label>
            <input type="number" name="dryMultiplier" value={editedItem.dryMultiplier} onChange={handleChange} readOnly />
          </div>
        )}
        {item.dryBase && (
          <div>
            <label>Dry Base:</label>
            <input type="number" name="dryBase" value={editedItem.dryBase} onChange={handleChange} readOnly />
          </div>
        )}
        {item.dryAllowedLength && (
          <div>
            <label>Dry Allowed Length:</label>
            <input type="number" name="dryAllowedLength" value={editedItem.dryAllowedLength} onChange={handleChange} readOnly />
          </div>
        )}
        {item.dryPenaltyLastN && (
          <div>
            <label>Dry Penalty Last N:</label>
            <input type="number" name="dryPenaltyLastN" value={editedItem.dryPenaltyLastN} onChange={handleChange} readOnly />
          </div>
        )}
        {item.repeatPenalty && (
          <div>
            <label>Repeat Penalty:</label>
            <input type="number" name="repeatPenalty" value={editedItem.repeatPenalty} onChange={handleChange} readOnly />
          </div>
        )}
        {item.repeatLastN && (
          <div>
            <label>Repeat Last N:</label>
            <input type="number" name="repeatLastN" value={editedItem.repeatLastN} onChange={handleChange} readOnly />
          </div>
        )}
        {item.topK && (
          <div>
            <label>Top K:</label>
            <input type="number" name="topK" value={editedItem.topK} onChange={handleChange} readOnly />
          </div>
        )}
        {item.topP && (
          <div>
            <label>Top P:</label>
            <input type="number" name="topP" value={editedItem.topP} onChange={handleChange} readOnly />
          </div>
        )}
        {item.minP && (
          <div>
            <label>Min P:</label>
            <input type="number" name="minP" value={editedItem.minP} onChange={handleChange} readOnly />
          </div>
        )}
        {item.topA && (
          <div>
            <label>Top A:</label>
            <input type="number" name="topA" value={editedItem.topA} onChange={handleChange} readOnly />
          </div>
        )}
        {item.xtcThreshold && (
          <div>
            <label>XTC Threshold:</label>
            <input type="number" name="xtcThreshold" value={editedItem.xtcThreshold} onChange={handleChange} readOnly />
          </div>
        )}
        {item.xtcProbability && (
          <div>
            <label>XTC Probability:</label>
            <input type="number" name="xtcProbability" value={editedItem.xtcProbability} onChange={handleChange} readOnly />
          </div>
        )}
        {item.temperature && (
          <div>
            <label>Temperature:</label>
            <input type="number" name="temperature" value={editedItem.temperature} onChange={handleChange} readOnly />
          </div>
        )}
        {categories && (
          <div>
            <label>Category:</label>
            <select name="category" value={selectedCategory} onChange={handleCategoryChange}>
              <option value="">Select a category</option>
              {categories.map((category, index) => (
                <option key={index} value={category}>{category}</option>
              ))}
            </select>
          </div>
        )}
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
