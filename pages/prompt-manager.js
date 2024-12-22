import { useState } from 'react';
import Layout from '../components/Layout';

export default function PromptManager() {
  const [prompts, setPrompts] = useState([]);
  const [formData, setFormData] = useState({
    content: '',
    solution: '',
    profile: '',
    category: ''
  });

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
          <input type="text" name="profile" value={formData.profile} onChange={handleChange} />
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
