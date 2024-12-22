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
      <h1 className="text-3xl font-bold mb-4">Profile Manager</h1>
      <form onSubmit={handleSubmit} className="bg-mystic-secondary p-4 rounded shadow-lg">
        <div className="mb-4">
          <label className="font-semibold">Name:</label>
          <input type="text" name="name" value={formData.name} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">System Prompt:</label>
          <textarea name="systemPrompt" value={formData.systemPrompt} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Dry Multiplier:</label>
          <input type="number" name="dryMultiplier" value={formData.dryMultiplier} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Dry Base:</label>
          <input type="number" name="dryBase" value={formData.dryBase} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Dry Allowed Length:</label>
          <input type="number" name="dryAllowedLength" value={formData.dryAllowedLength} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Dry Penalty Last N:</label>
          <input type="number" name="dryPenaltyLastN" value={formData.dryPenaltyLastN} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Repeat Penalty:</label>
          <input type="number" name="repeatPenalty" value={formData.repeatPenalty} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Repeat Last N:</label>
          <input type="number" name="repeatLastN" value={formData.repeatLastN} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Top K:</label>
          <input type="number" name="topK" value={formData.topK} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Top P:</label>
          <input type="number" name="topP" value={formData.topP} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Min P:</label>
          <input type="number" name="minP" value={formData.minP} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Top A:</label>
          <input type="number" name="topA" value={formData.topA} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">XTC Threshold:</label>
          <input type="number" name="xtcThreshold" value={formData.xtcThreshold} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">XTC Probability:</label>
          <input type="number" name="xtcProbability" value={formData.xtcProbability} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <div className="mb-4">
          <label className="font-semibold">Temperature:</label>
          <input type="number" name="temperature" value={formData.temperature} onChange={handleChange} className="bg-mystic-primary text-mystic-text p-2 rounded w-full" />
        </div>
        <button type="submit" className="bg-mystic-highlight text-mystic-primary px-4 py-2 rounded hover:bg-mystic-accent">Add Profile</button>
      </form>
      <ul className="list-disc pl-4">
        {profiles.map((profile, index) => (
          <li key={index} className="cursor-pointer hover:text-mystic-highlight" onClick={() => handleProfileClick(profile)}>
            {profile.name}
          </li>
        ))}
      </ul>
      <DetailPopup item={selectedProfile} onClose={handleClosePopup} />
    </Layout>
  );
}
