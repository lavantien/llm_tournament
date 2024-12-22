import React from 'react';
import styles from '../styles/DetailPopup.module.css';

const DetailPopup = ({ item, onClose }) => {
  if (!item) return null;

  return (
    <div className={styles.popupOverlay}>
      <div className={styles.popupContent}>
        <button className={styles.closeButton} onClick={onClose}>Close</button>
        <h2>Details</h2>
        {item.name && (
          <div>
            <label>Name:</label>
            <input type="text" value={item.name} readOnly />
          </div>
        )}
        {item.path && (
          <div>
            <label>Path:</label>
            <input type="text" value={item.path} readOnly />
          </div>
        )}
        {item.gpuLayers !== undefined && (
          <div>
            <label>GPU Layers:</label>
            <input type="number" value={item.gpuLayers} readOnly />
          </div>
        )}
        {item.ctxSize !== undefined && (
          <div>
            <label>Context Size:</label>
            <input type="number" value={item.ctxSize} readOnly />
          </div>
        )}
        {item.batchSize !== undefined && (
          <div>
            <label>Batch Size:</label>
            <input type="number" value={item.batchSize} readOnly />
          </div>
        )}
        {item.threads !== undefined && (
          <div>
            <label>Threads:</label>
            <input type="number" value={item.threads} readOnly />
          </div>
        )}
        {item.keep !== undefined && (
          <div>
            <label>Keep:</label>
            <input type="number" value={item.keep} readOnly />
          </div>
        )}
        {item.predict !== undefined && (
          <div>
            <label>Predict:</label>
            <input type="number" value={item.predict} readOnly />
          </div>
        )}
        {item.flashAttn !== undefined && (
          <div>
            <label>Flash Attn:</label>
            <input type="checkbox" checked={item.flashAttn} readOnly />
          </div>
        )}
        {item.mlock !== undefined && (
          <div>
            <label>Mlock:</label>
            <input type="checkbox" checked={item.mlock} readOnly />
          </div>
        )}
        {item.cacheTypeK && (
          <div>
            <label>Cache Type K:</label>
            <input type="text" value={item.cacheTypeK} readOnly />
          </div>
        )}
        {item.cacheTypeV && (
          <div>
            <label>Cache Type V:</label>
            <input type="text" value={item.cacheTypeV} readOnly />
          </div>
        )}
        {item.content && (
          <div>
            <label>Content:</label>
            <textarea value={item.content} readOnly />
          </div>
        )}
        {item.solution && (
          <div>
            <label>Solution:</label>
            <textarea value={item.solution} readOnly />
          </div>
        )}
        {item.profile && (
          <div>
            <label>Profile:</label>
            <input type="text" value={item.profile} readOnly />
          </div>
        )}
        {item.category && (
          <div>
            <label>Category:</label>
            <input type="text" value={item.category} readOnly />
          </div>
        )}
        {item.systemPrompt && (
          <div>
            <label>System Prompt:</label>
            <textarea value={item.systemPrompt} readOnly />
          </div>
        )}
        {item.dryMultiplier !== undefined && (
          <div>
            <label>Dry Multiplier:</label>
            <input type="number" value={item.dryMultiplier} readOnly />
          </div>
        )}
        {item.dryBase !== undefined && (
          <div>
            <label>Dry Base:</label>
            <input type="number" value={item.dryBase} readOnly />
          </div>
        )}
        {item.dryAllowedLength !== undefined && (
          <div>
            <label>Dry Allowed Length:</label>
            <input type="number" value={item.dryAllowedLength} readOnly />
          </div>
        )}
        {item.dryPenaltyLastN !== undefined && (
          <div>
            <label>Dry Penalty Last N:</label>
            <input type="number" value={item.dryPenaltyLastN} readOnly />
          </div>
        )}
        {item.repeatPenalty !== undefined && (
          <div>
            <label>Repeat Penalty:</label>
            <input type="number" value={item.repeatPenalty} readOnly />
          </div>
        )}
        {item.repeatLastN !== undefined && (
          <div>
            <label>Repeat Last N:</label>
            <input type="number" value={item.repeatLastN} readOnly />
          </div>
        )}
        {item.topK !== undefined && (
          <div>
            <label>Top K:</label>
            <input type="number" value={item.topK} readOnly />
          </div>
        )}
        {item.topP !== undefined && (
          <div>
            <label>Top P:</label>
            <input type="number" value={item.topP} readOnly />
          </div>
        )}
        {item.minP !== undefined && (
          <div>
            <label>Min P:</label>
            <input type="number" value={item.minP} readOnly />
          </div>
        )}
        {item.topA !== undefined && (
          <div>
            <label>Top A:</label>
            <input type="number" value={item.topA} readOnly />
          </div>
        )}
        {item.xtcThreshold !== undefined && (
          <div>
            <label>XTC Threshold:</label>
            <input type="number" value={item.xtcThreshold} readOnly />
          </div>
        )}
        {item.xtcProbability !== undefined && (
          <div>
            <label>XTC Probability:</label>
            <input type="number" value={item.xtcProbability} readOnly />
          </div>
        )}
        {item.temperature !== undefined && (
          <div>
            <label>Temperature:</label>
            <input type="number" value={item.temperature} readOnly />
          </div>
        )}
      </div>
    </div>
  );
};

export default DetailPopup;
