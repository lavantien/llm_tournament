import React from 'react';
import styles from '../styles/DetailPopup.module.css';

const DetailPopup = ({ item, onClose }) => {
  if (!item) return null;

  return (
    <div className={`fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 ${styles.popupOverlay}`}>
      <div className="bg-mystic-secondary p-6 rounded shadow-lg text-mystic-text">
        <button className="absolute top-0 right-0 mt-2 mr-2 text-mystic-highlight hover:text-mystic-accent" onClick={onClose}>
          &times;
        </button>
        <h2 className="text-2xl font-bold mb-4">Details</h2>
        {item.name && (
          <div className="mb-2">
            <label className="font-semibold">Name:</label>
            <p>{item.name}</p>
          </div>
        )}
        {item.path && (
          <div className="mb-2">
            <label className="font-semibold">Path:</label>
            <p>{item.path}</p>
          </div>
        )}
        {item.gpuLayers !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">GPU Layers:</label>
            <p>{item.gpuLayers}</p>
          </div>
        )}
        {item.ctxSize !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Context Size:</label>
            <p>{item.ctxSize}</p>
          </div>
        )}
        {item.batchSize !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Batch Size:</label>
            <p>{item.batchSize}</p>
          </div>
        )}
        {item.threads !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Threads:</label>
            <p>{item.threads}</p>
          </div>
        )}
        {item.keep !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Keep:</label>
            <p>{item.keep}</p>
          </div>
        )}
        {item.predict !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Predict:</label>
            <p>{item.predict}</p>
          </div>
        )}
        {item.flashAttn !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Flash Attn:</label>
            <p>{item.flashAttn ? 'True' : 'False'}</p>
          </div>
        )}
        {item.mlock !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Mlock:</label>
            <p>{item.mlock ? 'True' : 'False'}</p>
          </div>
        )}
        {item.cacheTypeK && (
          <div className="mb-2">
            <label className="font-semibold">Cache Type K:</label>
            <p>{item.cacheTypeK}</p>
          </div>
        )}
        {item.cacheTypeV && (
          <div className="mb-2">
            <label className="font-semibold">Cache Type V:</label>
            <p>{item.cacheTypeV}</p>
          </div>
        )}
        {item.content && (
          <div className="mb-2">
            <label className="font-semibold">Content:</label>
            <p>{item.content}</p>
          </div>
        )}
        {item.solution && (
          <div className="mb-2">
            <label className="font-semibold">Solution:</label>
            <p>{item.solution}</p>
          </div>
        )}
        {item.profile && (
          <div className="mb-2">
            <label className="font-semibold">Profile:</label>
            <p>{item.profile}</p>
          </div>
        )}
        {item.category && (
          <div className="mb-2">
            <label className="font-semibold">Category:</label>
            <p>{item.category}</p>
          </div>
        )}
        {item.systemPrompt && (
          <div className="mb-2">
            <label className="font-semibold">System Prompt:</label>
            <p>{item.systemPrompt}</p>
          </div>
        )}
        {item.dryMultiplier !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Dry Multiplier:</label>
            <p>{item.dryMultiplier}</p>
          </div>
        )}
        {item.dryBase !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Dry Base:</label>
            <p>{item.dryBase}</p>
          </div>
        )}
        {item.dryAllowedLength !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Dry Allowed Length:</label>
            <p>{item.dryAllowedLength}</p>
          </div>
        )}
        {item.dryPenaltyLastN !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Dry Penalty Last N:</label>
            <p>{item.dryPenaltyLastN}</p>
          </div>
        )}
        {item.repeatPenalty !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Repeat Penalty:</label>
            <p>{item.repeatPenalty}</p>
          </div>
        )}
        {item.repeatLastN !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Repeat Last N:</label>
            <p>{item.repeatLastN}</p>
          </div>
        )}
        {item.topK !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Top K:</label>
            <p>{item.topK}</p>
          </div>
        )}
        {item.topP !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Top P:</label>
            <p>{item.topP}</p>
          </div>
        )}
        {item.minP !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Min P:</label>
            <p>{item.minP}</p>
          </div>
        )}
        {item.topA !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Top A:</label>
            <p>{item.topA}</p>
          </div>
        )}
        {item.xtcThreshold !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">XTC Threshold:</label>
            <p>{item.xtcThreshold}</p>
          </div>
        )}
        {item.xtcProbability !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">XTC Probability:</label>
            <p>{item.xtcProbability}</p>
          </div>
        )}
        {item.temperature !== undefined && (
          <div className="mb-2">
            <label className="font-semibold">Temperature:</label>
            <p>{item.temperature}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default DetailPopup;
