import React, { useState } from 'react';
import styles from '../styles/DetailPopup.module.css';

const DetailPopup = ({ item, onClose, onSave }) => {
  const [editedItem, setEditedItem] = useState({ ...item });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setEditedItem({
      ...editedItem,
      [name]: value
    });
  };

  const handleSave = () => {
    onSave(editedItem);
    onClose();
  };

  if (!item) return null;

  return (
    <div className={styles.popupOverlay}>
      <div className={styles.popupContent}>
        <button className={styles.closeButton} onClick={onClose}>Close</button>
        <h2>Details</h2>
        <div>
          <label>Name:</label>
          <input type="text" name="name" value={editedItem.name} onChange={handleChange} readOnly />
        </div>
        {Object.keys(editedItem).filter(key => key.startsWith('category')).map((key, index) => (
          <div key={index}>
            <label>{key}:</label>
            <input type="number" name={key} value={editedItem[key]} onChange={handleChange} />
          </div>
        ))}
        <div>
          <label>Overall Score:</label>
          <input type="number" name="overall" value={editedItem.overall} onChange={handleChange} readOnly />
        </div>
        <button onClick={handleSave}>Save</button>
      </div>
    </div>
  );
};

export default DetailPopup;
