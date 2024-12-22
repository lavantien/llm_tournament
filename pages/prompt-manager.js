import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import { getProfiles } from './profile-manager';
import DetailPopup from '../components/DetailPopup';

export default function PromptManager() {
  const [prompts, setPrompts] = useState([
    {
      content: 'Prompt 1 content',
      solution: 'Prompt 1 solution',
      profile: 'Profile 1',
      category: 'Category 1'
    },
    {
      content: 'Prompt 2 content',
      solution: 'Prompt 2 solution',
      profile: 'Profile 2',
      category: 'Category 2'
    },
    {
      content: 'Prompt 3 content',
      solution: 'Prompt 3 solution',
      profile: 'Profile 3',
      category: 'Category 3'
    }
    // Add more mock data as needed
  ]);
  const [formData, setFormData] = useState({
    content: '',
    solution: '',
    profile: '',
    category: ''
  });
  const [profiles, setProfiles] = useState([]);
  const [selectedPrompt, setSelectedPrompt] = useState(null);

  useEffect(() => {
    // Fetch profiles from the ProfileManager
    const profilesData = getProfiles();
    setProfiles(profilesData);
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
    // Handle form submission (e.g., add prompt to the database)
    setPrompts([...prompts, formData]);
    setFormData({
      content: '',
      solution: '',
      profile: '',
      category: ''
    });
  };

  const handlePromptClick = (prompt) => {
    setSelectedPrompt(prompt);
  };

  const handleClosePopup = () => {
    setSelectedPrompt(null);
  };

  return (
    <Layout>
      <h1>Prompt Manager</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Content:</label>
          <textarea name="content" value={formData.content} onChange={handleChange} />
        </div>
        <div>
          <label>Solution:</label>
          <textarea name="solution" value={formData.solution} onChange={handleChange} />
        </div>
        <div>
          <label>Profile:</label>
          <select name="profile" value={formData.profile} onChange={handleChange}>
            <option value="">Select a profile</option>
            {profiles.map((profile, index) => (
              <option key={index} value={profile.name}>{profile.name}</option>
            ))}
          </select>
        </div>
        <div>
          <label>Category:</label>
          <input type="text" name="category" value={formData.category} onChange={handleChange} />
        </div>
        <button type="submit">Add Prompt</button>
      </form>
      <ul>
        {prompts.map((prompt, index) => (
          <li key={index} onClick={() => handlePromptClick(prompt)}>
            {prompt.content}
          </li>
        ))}
      </ul>
      <DetailPopup item={selectedPrompt} onClose={handleClosePopup} />
    </Layout>
  );
}
