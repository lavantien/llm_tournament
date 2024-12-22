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
            <p>{item.name}</p>
          </div>
        )}
        {item.path && (
          <div>
            <label>Path:</label>
            <p>{item.path}</p>
          </div>
        )}
        {item.gpuLayers !== undefined && (
          <div>
            <label>GPU Layers:</label>
            <p>{item.gpuLayers}</p>
          </div>
        )}
        {item.ctxSize !== undefined && (
          <div>
            <label>Context Size:</label>
            <p>{item.ctxSize}</p>
          </div>
        )}
        {item.batchSize !== undefined && (
          <div>
            <label>Batch Size:</label>
            <p>{item.batchSize}</p>
          </div>
        )}
        {item.threads !== undefined && (
          <div>
            <label>Threads:</label>
            <p>{item.threads}</p>
          </div>
        )}
        {item.keep !== undefined && (
          <div>
            <label>Keep:</label>
            <p>{item.keep}</p>
          </div>
        )}
        {item.predict !== undefined && (
          <div>
            <label>Predict:</label>
            <p>{item.predict}</p>
          </div>
        )}
        {item.flashAttn !== undefined && (
          <div>
            <label>Flash Attn:</label>
            <p>{item.flashAttn ? 'True' : 'False'}</p>
          </div>
        )}
        {item.mlock !== undefined && (
          <div>
            <label>Mlock:</label>
            <p>{item.mlock ? 'True' : 'False'}</p>
          </div>
        )}
        {item.cacheTypeK && (
          <div>
            <label>Cache Type K:</label>
            <p>{item.cacheTypeK}</p>
          </div>
        )}
        {item.cacheTypeV && (
          <div>
            <label>Cache Type V:</label>
            <p>{item.cacheTypeV}</p>
          </div>
        )}
        {item.content && (
          <div>
            <label>Content:</label>
            <p>{item.content}</p>
          </div>
        )}
        {item.solution && (
          <div>
            <label>Solution:</label>
            <p>{item.solution}</p>
          </div>
        )}
        {item.profile && (
          <div>
            <label>Profile:</label>
            <p>{item.profile}</p>
          </div>
        )}
        {item.category && (
          <div>
            <label>Category:</label>
            <p>{item.category}</p>
          </div>
        )}
        {item.systemPrompt && (
          <div>
            <label>System Prompt:</label>
            <p>{item.systemPrompt}</p>
          </div>
        )}
        {item.dryMultiplier !== undefined && (
          <div>
            <label>Dry Multiplier:</label>
            <p>{item.dryMultiplier}</p>
          </div>
        )}
        {item.dryBase !== undefined && (
          <div>
            <label>Dry Base:</label>
            <p>{item.dryBase}</p>
          </div>
        )}
        {item.dryAllowedLength !== undefined && (
          <div>
            <label>Dry Allowed Length:</label>
            <p>{item.dryAllowedLength}</p>
          </div>
        )}
        {item.dryPenaltyLastN !== undefined && (
          <div>
            <label>Dry Penalty Last N:</label>
            <p>{item.dryPenaltyLastN}</p>
          </div>
        )}
        {item.repeatPenalty !== undefined && (
          <div>
            <label>Repeat Penalty:</label>
            <p>{item.repeatPenalty}</p>
          </div>
        )}
        {item.repeatLastN !== undefined && (
          <div>
            <label>Repeat Last N:</label>
            <p>{item.repeatLastN}</p>
          </div>
        )}
        {item.topK !== undefined && (
          <div>
            <label>Top K:</label>
            <p>{item.topK}</p>
          </div>
        )}
        {item.topP !== undefined && (
          <div>
            <label>Top P:</label>
            <p>{item.topP}</p>
          </div>
        )}
        {item.minP !== undefined && (
          <div>
            <label>Min P:</label>
            <p>{item.minP}</p>
          </div>
        )}
        {item.topA !== undefined && (
          <div>
            <label>Top A:</label>
            <p>{item.topA}</p>
          </div>
        )}
        {item.xtcThreshold !== undefined && (
          <div>
            <label>XTC Threshold:</label>
            <p>{item.xtcThreshold}</p>
          </div>
        )}
        {item.xtcProbability !== undefined && (
          <div>
            <label>XTC Probability:</label>
            <p>{item.xtcProbability}</p>
          </div>
        )}
        {item.temperature !== undefined && (
          <div>
            <label>Temperature:</label>
            <p>{item.temperature}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default DetailPopup;
