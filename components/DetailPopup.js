import React from 'react';

const DetailPopup = ({ item, onClose }) => {
  if (!item) return null;

  return (
    <div id="myModal" className="modal">
      <div className="modal-content">
        <span className="close" onClick={onClose}>&times;</span>
        <h2>Details</h2>
        {item.name && (
          <div className="form-group">
            <label className="form-label">Name:</label>
            <p>{item.name}</p>
          </div>
        )}
        {item.path && (
          <div className="form-group">
            <label className="form-label">Path:</label>
            <p>{item.path}</p>
          </div>
        )}
        {item.gpuLayers !== undefined && (
          <div className="form-group">
            <label className="form-label">GPU Layers:</label>
            <p>{item.gpuLayers}</p>
          </div>
        )}
        {item.ctxSize !== undefined && (
          <div className="form-group">
            <label className="form-label">Context Size:</label>
            <p>{item.ctxSize}</p>
          </div>
        )}
        {item.batchSize !== undefined && (
          <div className="form-group">
            <label className="form-label">Batch Size:</label>
            <p>{item.batchSize}</p>
          </div>
        )}
        {item.threads !== undefined && (
          <div className="form-group">
            <label className="form-label">Threads:</label>
            <p>{item.threads}</p>
          </div>
        )}
        {item.keep !== undefined && (
          <div className="form-group">
            <label className="form-label">Keep:</label>
            <p>{item.keep}</p>
          </div>
        )}
        {item.predict !== undefined && (
          <div className="form-group">
            <label className="form-label">Predict:</label>
            <p>{item.predict}</p>
          </div>
        )}
        {item.flashAttn !== undefined && (
          <div className="form-group">
            <label className="form-label">Flash Attn:</label>
            <p>{item.flashAttn ? 'True' : 'False'}</p>
          </div>
        )}
        {item.mlock !== undefined && (
          <div className="form-group">
            <label className="form-label">Mlock:</label>
            <p>{item.mlock ? 'True' : 'False'}</p>
          </div>
        )}
        {item.cacheTypeK && (
          <div className="form-group">
            <label className="form-label">Cache Type K:</label>
            <p>{item.cacheTypeK}</p>
          </div>
        )}
        {item.cacheTypeV && (
          <div className="form-group">
            <label className="form-label">Cache Type V:</label>
            <p>{item.cacheTypeV}</p>
          </div>
        )}
        {item.content && (
          <div className="form-group">
            <label className="form-label">Content:</label>
            <p>{item.content}</p>
          </div>
        )}
        {item.solution && (
          <div className="form-group">
            <label className="form-label">Solution:</label>
            <p>{item.solution}</p>
          </div>
        )}
        {item.profile && (
          <div className="form-group">
            <label className="form-label">Profile:</label>
            <p>{item.profile}</p>
          </div>
        )}
        {item.category && (
          <div className="form-group">
            <label className="form-label">Category:</label>
            <p>{item.category}</p>
          </div>
        )}
        {item.systemPrompt && (
          <div className="form-group">
            <label className="form-label">System Prompt:</label>
            <p>{item.systemPrompt}</p>
          </div>
        )}
        {item.dryMultiplier !== undefined && (
          <div className="form-group">
            <label className="form-label">Dry Multiplier:</label>
            <p>{item.dryMultiplier}</p>
          </div>
        )}
        {item.dryBase !== undefined && (
          <div className="form-group">
            <label className="form-label">Dry Base:</label>
            <p>{item.dryBase}</p>
          </div>
        )}
        {item.dryAllowedLength !== undefined && (
          <div className="form-group">
            <label className="form-label">Dry Allowed Length:</label>
            <p>{item.dryAllowedLength}</p>
          </div>
        )}
        {item.dryPenaltyLastN !== undefined && (
          <div className="form-group">
            <label className="form-label">Dry Penalty Last N:</label>
            <p>{item.dryPenaltyLastN}</p>
          </div>
        )}
        {item.repeatPenalty !== undefined && (
          <div className="form-group">
            <label className="form-label">Repeat Penalty:</label>
            <p>{item.repeatPenalty}</p>
          </div>
        )}
        {item.repeatLastN !== undefined && (
          <div className="form-group">
            <label className="form-label">Repeat Last N:</label>
            <p>{item.repeatLastN}</p>
          </div>
        )}
        {item.topK !== undefined && (
          <div className="form-group">
            <label className="form-label">Top K:</label>
            <p>{item.topK}</p>
          </div>
        )}
        {item.topP !== undefined && (
          <div className="form-group">
            <label className="form-label">Top P:</label>
            <p>{item.topP}</p>
          </div>
        )}
        {item.minP !== undefined && (
          <div className="form-group">
            <label className="form-label">Min P:</label>
            <p>{item.minP}</p>
          </div>
        )}
        {item.topA !== undefined && (
          <div className="form-group">
            <label className="form-label">Top A:</label>
            <p>{item.topA}</p>
          </div>
        )}
        {item.xtcThreshold !== undefined && (
          <div className="form-group">
            <label className="form-label">XTC Threshold:</label>
            <p>{item.xtcThreshold}</p>
          </div>
        )}
        {item.xtcProbability !== undefined && (
          <div className="form-group">
            <label className="form-label">XTC Probability:</label>
            <p>{item.xtcProbability}</p>
          </div>
        )}
        {item.temperature !== undefined && (
          <div className="form-group">
            <label className="form-label">Temperature:</label>
            <p>{item.temperature}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default DetailPopup;
