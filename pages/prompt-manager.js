import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import { getProfiles } from './profile-manager';

export default function PromptManager() {
  const [prompts, setPrompts] = useState([]);
  const [formData, setFormData] = useState({
    content: '',
    solution: '',
    profile: '',
    category: ''
  });
  const [profiles, setProfiles] = useState([]);

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
          <li key={index}>{prompt.content}</li>
        ))}
      </ul>
    </Layout>
  );
}
