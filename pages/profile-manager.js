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
    const fetchProfiles = async () => {
      try {
        const response = await fetch('/api/profiles');
        const data = await response.json();
        setProfiles(data);
        console.log('Profiles set:', data);
      } catch (error) {
        console.error('Error fetching profiles:', error);
      }
    };

    fetchProfiles();

    // Listen for dataWiped event
    const handleDataWiped = () => {
      setProfiles([]);
      console.log('Data wiped, profiles reset');
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
    console.log('Form data set:', { ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission (e.g., add profile to the database)
    setProfiles([...profiles, formData]);
    console.log('Profiles set:', [...profiles, formData]);
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
    console.log('Form data reset');
  };

  const handleProfileClick = (profile) => {
    setSelectedProfile(profile);
    console.log('Selected profile set:', profile);
  };

  const handleClosePopup = () => {
    setSelectedProfile(null);
    console.log('Selected profile reset');
  };

  return (
    <Layout>
      <h1 className="h3 mb-4">Profile Manager</h1>
      <form onSubmit={handleSubmit} className="bg-secondary p-3 rounded shadow">
        <div className="mb-3">
          <label className="form-label">Name:</label>
          <input type="text" name="name" value={formData.name} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">System Prompt:</label>
          <textarea name="systemPrompt" value={formData.systemPrompt} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Dry Multiplier:</label>
          <input type="number" name="dryMultiplier" value={formData.dryMultiplier} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Dry Base:</label>
          <input type="number" name="dryBase" value={formData.dryBase} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Dry Allowed Length:</label>
          <input type="number" name="dryAllowedLength" value={formData.dryAllowedLength} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Dry Penalty Last N:</label>
          <input type="number" name="dryPenaltyLastN" value={formData.dryPenaltyLastN} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Repeat Penalty:</label>
          <input type="number" name="repeatPenalty" value={formData.repeatPenalty} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Repeat Last N:</label>
          <input type="number" name="repeatLastN" value={formData.repeatLastN} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Top K:</label>
          <input type="number" name="topK" value={formData.topK} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Top P:</label>
          <input type="number" name="topP" value={formData.topP} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Min P:</label>
          <input type="number" name="minP" value={formData.minP} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Top A:</label>
          <input type="number" name="topA" value={formData.topA} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">XTC Threshold:</label>
          <input type="number" name="xtcThreshold" value={formData.xtcThreshold} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">XTC Probability:</label>
          <input type="number" name="xtcProbability" value={formData.xtcProbability} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Temperature:</label>
          <input type="number" name="temperature" value={formData.temperature} onChange={handleChange} className="form-control" />
        </div>
        <button type="submit" className="btn btn-primary">Add Profile</button>
      </form>
      <ul className="list-group">
        {profiles.map((profile, index) => (
          <li key={index} className="list-group-item list-group-item-action" onClick={() => handleProfileClick(profile)}>
            {profile.name}
          </li>
        ))}
      </ul>
      <DetailPopup item={selectedProfile} onClose={handleClosePopup} />
    </Layout>
  );
}
