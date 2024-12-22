import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';

export default function ProfileManager() {
  const [profiles, setProfiles] = useState([
    {
      name: 'Profile 1',
      systemPrompt: 'System prompt 1',
      dryMultiplier: 0.8,
      dryBase: 1.75,
      dryAllowedLength: 2,
      dryPenaltyLastN: 512,
      repeatPenalty: 1.02,
      repeatLastN: 512,
      topK: 0,
      topP: 1,
      minP: 0.02,
      topA: 0.12,
      xtcThreshold: 0.1,
      xtcProbability: 0.5,
      temperature: 1
    },
    {
      name: 'Profile 2',
      systemPrompt: 'System prompt 2',
      dryMultiplier: 0.9,
      dryBase: 1.8,
      dryAllowedLength: 3,
      dryPenaltyLastN: 512,
      repeatPenalty: 1.03,
      repeatLastN: 512,
      topK: 1,
      topP: 1.1,
      minP: 0.03,
      topA: 0.13,
      xtcThreshold: 0.2,
      xtcProbability: 0.6,
      temperature: 1.1
    },
    {
      name: 'Profile 3',
      systemPrompt: 'System prompt 3',
      dryMultiplier: 1.0,
      dryBase: 1.9,
      dryAllowedLength: 4,
      dryPenaltyLastN: 512,
      repeatPenalty: 1.04,
      repeatLastN: 512,
      topK: 2,
      topP: 1.2,
      minP: 0.04,
      topA: 0.14,
      xtcThreshold: 0.3,
      xtcProbability: 0.7,
      temperature: 1.2
    }
    // Add more mock data as needed
  ]);
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
    // Fetch profiles from the database or any other source
    // For now, we'll use dummy data
    setProfiles([
      { name: 'Profile 1', systemPrompt: 'Prompt 1' },
      { name: 'Profile 2', systemPrompt: 'Prompt 2' }
    ]);
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

export function getProfiles() {
  // This function will be used to get the list of profiles
  return [
    { name: 'Profile 1', systemPrompt: 'Prompt 1' },
    { name: 'Profile 2', systemPrompt: 'Prompt 2' },
    { name: 'Profile 3', systemPrompt: 'Prompt 3' }
  ];
}
