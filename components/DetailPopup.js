import React from 'react';
import { Modal, Button } from 'react-bootstrap';

const DetailPopup = ({ item, onClose }) => {
  if (!item) return null;

  return (
    <Modal show onHide={onClose} centered>
      <Modal.Header closeButton>
        <Modal.Title>Details</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        {item.name && (
          <div className="mb-2">
            <label className="form-label">Name:</label>
            <p>{item.name}</p>
          </div>
        )}
        {item.path && (
          <div className="mb-2">
            <label className="form-label">Path:</label>
            <p>{item.path}</p>
          </div>
        )}
        {item.gpuLayers !== undefined && (
          <div className="mb-2">
            <label className="form-label">GPU Layers:</label>
            <p>{item.gpuLayers}</p>
          </div>
        )}
        {item.ctxSize !== undefined && (
          <div className="mb-2">
            <label className="form-label">Context Size:</label>
            <p>{item.ctxSize}</p>
          </div>
        )}
        {item.batchSize !== undefined && (
          <div className="mb-2">
            <label className="form-label">Batch Size:</label>
            <p>{item.batchSize}</p>
          </div>
        )}
        {item.threads !== undefined && (
          <div className="mb-2">
            <label className="form-label">Threads:</label>
            <p>{item.threads}</p>
          </div>
        )}
        {item.keep !== undefined && (
          <div className="mb-2">
            <label className="form-label">Keep:</label>
            <p>{item.keep}</p>
          </div>
        )}
        {item.predict !== undefined && (
          <div className="mb-2">
            <label className="form-label">Predict:</label>
            <p>{item.predict}</p>
          </div>
        )}
        {item.flashAttn !== undefined && (
          <div className="mb-2">
            <label className="form-label">Flash Attn:</label>
            <p>{item.flashAttn ? 'True' : 'False'}</p>
          </div>
        )}
        {item.mlock !== undefined && (
          <div className="mb-2">
            <label className="form-label">Mlock:</label>
            <p>{item.mlock ? 'True' : 'False'}</p>
          </div>
        )}
        {item.cacheTypeK && (
          <div className="mb-2">
            <label className="form-label">Cache Type K:</label>
            <p>{item.cacheTypeK}</p>
          </div>
        )}
        {item.cacheTypeV && (
          <div className="mb-2">
            <label className="form-label">Cache Type V:</label>
            <p>{item.cacheTypeV}</p>
          </div>
        )}
        {item.content && (
          <div className="mb-2">
            <label className="form-label">Content:</label>
            <p>{item.content}</p>
          </div>
        )}
        {item.solution && (
          <div className="mb-2">
            <label className="form-label">Solution:</label>
            <p>{item.solution}</p>
          </div>
        )}
        {item.profile && (
          <div className="mb-2">
            <label className="form-label">Profile:</label>
            <p>{item.profile}</p>
          </div>
        )}
        {item.category && (
          <div className="mb-2">
            <label className="form-label">Category:</label>
            <p>{item.category}</p>
          </div>
        )}
        {item.systemPrompt && (
          <div className="mb-2">
            <label className="form-label">System Prompt:</label>
            <p>{item.systemPrompt}</p>
          </div>
        )}
        {item.dryMultiplier !== undefined && (
          <div className="mb-2">
            <label className="form-label">Dry Multiplier:</label>
            <p>{item.dryMultiplier}</p>
          </div>
        )}
        {item.dryBase !== undefined && (
          <div className="mb-2">
            <label className="form-label">Dry Base:</label>
            <p>{item.dryBase}</p>
          </div>
        )}
        {item.dryAllowedLength !== undefined && (
          <div className="mb-2">
            <label className="form-label">Dry Allowed Length:</label>
            <p>{item.dryAllowedLength}</p>
          </div>
        )}
        {item.dryPenaltyLastN !== undefined && (
          <div className="mb-2">
            <label className="form-label">Dry Penalty Last N:</label>
            <p>{item.dryPenaltyLastN}</p>
          </div>
        )}
        {item.repeatPenalty !== undefined && (
          <div className="mb-2">
            <label className="form-label">Repeat Penalty:</label>
            <p>{item.repeatPenalty}</p>
          </div>
        )}
        {item.repeatLastN !== undefined && (
          <div className="mb-2">
            <label className="form-label">Repeat Last N:</label>
            <p>{item.repeatLastN}</p>
          </div>
        )}
        {item.topK !== undefined && (
          <div className="mb-2">
            <label className="form-label">Top K:</label>
            <p>{item.topK}</p>
          </div>
        )}
        {item.topP !== undefined && (
          <div className="mb-2">
            <label className="form-label">Top P:</label>
            <p>{item.topP}</p>
          </div>
        )}
        {item.minP !== undefined && (
          <div className="mb-2">
            <label className="form-label">Min P:</label>
            <p>{item.minP}</p>
          </div>
        )}
        {item.topA !== undefined && (
          <div className="mb-2">
            <label className="form-label">Top A:</label>
            <p>{item.topA}</p>
          </div>
        )}
        {item.xtcThreshold !== undefined && (
          <div className="mb-2">
            <label className="form-label">XTC Threshold:</label>
            <p>{item.xtcThreshold}</p>
          </div>
        )}
        {item.xtcProbability !== undefined && (
          <div className="mb-2">
            <label className="form-label">XTC Probability:</label>
            <p>{item.xtcProbability}</p>
          </div>
        )}
        {item.temperature !== undefined && (
          <div className="mb-2">
            <label className="form-label">Temperature:</label>
            <p>{item.temperature}</p>
          </div>
        )}
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={onClose}>
          Close
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default DetailPopup;
