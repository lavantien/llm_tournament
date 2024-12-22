import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';

export default function PromptManager() {
  const [prompts, setPrompts] = useState([]);
  const [profiles, setProfiles] = useState([]);
  const [categories, setCategories] = useState([]);
  const [formData, setFormData] = useState({
    content: '',
    solution: '',
    profile: '',
    category: ''
  });
  const [selectedPrompt, setSelectedPrompt] = useState(null);

  useEffect(() => {
    // Fetch prompts from the API route
    const fetchPrompts = async () => {
      try {
        const response = await fetch('/api/prompts');
        const data = await response.json();
        setPrompts(data);
        console.log('Prompts set:', data);
      } catch (error) {
        console.error('Error fetching prompts:', error);
      }
    };

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

    fetchPrompts();
    fetchProfiles();

    // Listen for dataWiped event
    const handleDataWiped = () => {
      setPrompts([]);
      setProfiles([]);
      setCategories([]);
      console.log('Data wiped, state reset');
    };

    window.addEventListener('dataWiped', handleDataWiped);

    return () => {
      window.removeEventListener('dataWiped', handleDataWiped);
    };
  }, []);

  useEffect(() => {
    // Fetch categories from the prompts
    const uniqueCategories = [...new Set(prompts.map(prompt => prompt.category))];
    setCategories(uniqueCategories);
    console.log('Categories set:', uniqueCategories);
  }, [prompts]);

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
    // Handle form submission (e.g., add prompt to the database)
    setPrompts([...prompts, formData]);
    console.log('Prompts set:', [...prompts, formData]);
    setFormData({
      content: '',
      solution: '',
      profile: '',
      category: ''
    });
    console.log('Form data reset');
  };

  const handlePromptClick = (prompt) => {
    setSelectedPrompt(prompt);
    console.log('Selected prompt set:', prompt);
  };

  const handleClosePopup = () => {
    setSelectedPrompt(null);
    console.log('Selected prompt reset');
  };

  return (
    <Layout>
      <h1 className="h3 mb-4">Prompt Manager</h1>
      <form onSubmit={handleSubmit} className="bg-secondary p-3 rounded shadow">
        <div className="mb-3">
          <label className="form-label">Content:</label>
          <textarea name="content" value={formData.content} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Solution:</label>
          <textarea name="solution" value={formData.solution} onChange={handleChange} className="form-control" />
        </div>
        <div className="mb-3">
          <label className="form-label">Profile:</label>
          <select name="profile" value={formData.profile} onChange={handleChange} className="form-select">
            <option value="">Select a profile</option>
            {profiles.map((profile, index) => (
              <option key={index} value={profile.name}>{profile.name}</option>
            ))}
          </select>
        </div>
        <div className="mb-3">
          <label className="form-label">Category:</label>
          <select name="category" value={formData.category} onChange={handleChange} className="form-select">
            <option value="">Select a category</option>
            {categories.map((category, index) => (
              <option key={index} value={category}>{category}</option>
            ))}
          </select>
        </div>
        <button type="submit" className="btn btn-primary">Add Prompt</button>
      </form>
      <ul className="list-group">
        {prompts.map((prompt, index) => (
          <li key={index} className="list-group-item list-group-item-action" onClick={() => handlePromptClick(prompt)}>
            {prompt.content}
          </li>
        ))}
      </ul>
      <DetailPopup item={selectedPrompt} onClose={handleClosePopup} />
    </Layout>
  );
}
