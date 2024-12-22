import React from 'react';
import styles from '../styles/DetailPopup.module.css';

const DetailPopup = ({ item, onClose }) => {
  if (!item) return null;

  return (
    <div className={styles.popupOverlay}>
      <div className={styles.popupContent}>
        <button className={styles.closeButton} onClick={onClose}>Close</button>
        <h2>Details</h2>
        <pre>{JSON.stringify(item, null, 2)}</pre>
      </div>
    </div>
  );
};

export default DetailPopup;
