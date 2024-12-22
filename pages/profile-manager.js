import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';

export default function ProfileManager() {
  const [profiles, setProfiles] = useState([]);
  const [formData, setFormData] = useState({
    name: '',
    systemPrompt: '',
    dryMultiplier: '',
    dryBase: '',
    dryAllowedLength: '',
    dryPenaltyLastN: '',
    repeatPenalty: '',
    repeatLastN: '',
    topK: '',
    topP: '',
    minP: '',
    topA: '',
    xtcThreshold: '',
    xtcProbability: '',
    temperature: ''
  });
  const [selectedProfile, setSelectedProfile] = useState(null);

  useEffect(() => {
    // Fetch profiles from the API route
    fetch('/api/mocks/profiles')
      .then(response => response.json())
      .then(data => setProfiles(data))
      .catch(error => console.error('Error fetching profiles:', error));

    // Listen for dataWiped event
    const handleDataWiped = () => {
      setProfiles([]);
    };

    window.addEventListener('dataWiped', handleDataWiped);

    return () => {
      window.removeEventListener('dataWiped', handleDataWiped);
    };
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission (e.g., add profile to the database)
    setProfiles([...profiles, formData]);
    setFormData({
      name: '',
      systemPrompt: '',
      dryMultiplier: '',
      dryBase: '',
      dryAllowedLength: '',
      dryPenaltyLastN: '',
      repeatPenalty: '',
      repeatLastN: '',
      topK: '',
      topP: '',
      minP: '',
      topA: '',
      xtcThreshold: '',
      xtcProbability: '',
      temperature: ''
    });
  };

  const handleProfileClick = (profile) => {
    setSelectedProfile(profile);
  };

  const handleClosePopup = () => {
    setSelectedProfile(null);
  };

  return (
    <Layout>
      <h1>Profile Manager</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name:</label>
          <input type="text" name="name" value={formData.name} onChange={handleChange} />
        </div>
        <div>
          <label>System Prompt:</label>
          <textarea name="systemPrompt" value={formData.systemPrompt} onChange={handleChange} />
        </div>
        <div>
          <label>Dry Multiplier:</label>
          <input type="number" name="dryMultiplier" value={formData.dryMultiplier} onChange={handleChange} />
        </div>
        <div>
          <label>Dry Base:</label>
          <input type="number" name="dryBase" value={formData.dryBase} onChange={handleChange} />
        </div>
        <div>
          <label>Dry Allowed Length:</label>
          <input type="number" name="dryAllowedLength" value={formData.dryAllowedLength} onChange={handleChange} />
        </div>
        <div>
          <label>Dry Penalty Last N:</label>
          <input type="number" name="dryPenaltyLastN" value={formData.dryPenaltyLastN} onChange={handleChange} />
        </div>
        <div>
          <label>Repeat Penalty:</label>
          <input type="number" name="repeatPenalty" value={formData.repeatPenalty} onChange={handleChange} />
        </div>
        <div>
          <label>Repeat Last N:</label>
          <input type="number" name="repeatLastN" value={formData.repeatLastN} onChange={handleChange} />
        </div>
        <div>
          <label>Top K:</label>
          <input type="number" name="topK" value={formData.topK} onChange={handleChange} />
        </div>
        <div>
          <label>Top P:</label>
          <input type="number" name="topP" value={formData.topP} onChange={handleChange} />
        </div>
        <div>
          <label>Min P:</label>
          <input type="number" name="minP" value={formData.minP} onChange={handleChange} />
        </div>
        <div>
          <label>Top A:</label>
          <input type="number" name="topA" value={formData.topA} onChange={handleChange} />
        </div>
        <div>
          <label>XTC Threshold:</label>
          <input type="number" name="xtcThreshold" value={formData.xtcThreshold} onChange={handleChange} />
        </div>
        <div>
          <label>XTC Probability:</label>
          <input type="number" name="xtcProbability" value={formData.xtcProbability} onChange={handleChange} />
        </div>
        <div>
          <label>Temperature:</label>
          <input type="number" name="temperature" value={formData.temperature} onChange={handleChange} />
        </div>
        <button type="submit">Add Profile</button>
      </form>
      <ul>
        {profiles.map((profile, index) => (
          <li key={index} onClick={() => handleProfileClick(profile)}>
            {profile.name}
          </li>
        ))}
      </ul>
      <DetailPopup item={selectedProfile} onClose={handleClosePopup} />
    </Layout>
  );
}
