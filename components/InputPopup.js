import React, { useState, useEffect } from 'react';
import styles from '../styles/DetailPopup.module.css';

const InputPopup = ({ item, onClose, onSave, categories, prompts, scores }) => {
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

  const filteredPrompts = prompts.filter(prompt => prompt.category === selectedCategory);

  const modelScores = scores.filter(score => score.modelId === item.id);

  return (
    <div className={`fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 ${styles.popupOverlay}`}>
      <div className="bg-mystic-secondary p-6 rounded shadow-lg text-mystic-text">
        <button className="absolute top-0 right-0 mt-2 mr-2 text-mystic-highlight hover:text-mystic-accent" onClick={onClose}>
          &times;
        </button>
        <h2 className="text-2xl font-bold mb-4">Edit Scores</h2>
        <div className="mb-4">
          <label className="font-semibold">Category:</label>
          <select value={selectedCategory} onChange={handleCategoryChange} className="bg-mystic-primary text-mystic-text p-2 rounded">
            <option value="">Select a category</option>
            {categories.map((category, index) => (
              <option key={index} value={category}>
                {category}
              </option>
            ))}
          </select>
        </div>
        {selectedCategory && (
          <div className="mb-4">
            <label className="font-semibold">Prompt:</label>
            <select value={selectedPrompt} onChange={handlePromptChange} className="bg-mystic-primary text-mystic-text p-2 rounded">
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
          <div className="mb-4">
            <label className="font-semibold">Attempts:</label>
            <input type="number" value={attempts} onChange={handleAttemptsChange} min="1" className="bg-mystic-primary text-mystic-text p-2 rounded" />
          </div>
        )}
        <button className="bg-mystic-highlight text-mystic-primary px-4 py-2 rounded hover:bg-mystic-accent" onClick={handleSave}>Save</button>
        <h3 className="text-xl font-bold mt-4 mb-2">Scores for {item.name}</h3>
        <table className="min-w-full bg-mystic-primary text-mystic-text">
          <thead>
            <tr>
              <th className="py-2 px-4 border-b">Prompt</th>
              <th className="py-2 px-4 border-b">Score</th>
            </tr>
          </thead>
          <tbody>
            {modelScores.map((score, index) => {
              const prompt = prompts.find(p => p.id === score.promptId);
              return (
                <tr key={index}>
                  <td className="py-2 px-4 border-b">{prompt ? prompt.content : 'Unknown Prompt'}</td>
                  <td className="py-2 px-4 border-b">{score.score}</td>
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
