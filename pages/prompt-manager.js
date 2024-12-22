import { useState, useEffect } from 'react';
import Layout from '../components/Layout';
import DetailPopup from '../components/DetailPopup';
import { Form, Button, Container, ListGroup } from 'react-bootstrap';

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
      <Container>
        <h1 className="h3 mb-4">Prompt Manager</h1>
        <Form onSubmit={handleSubmit} className="bg-secondary p-3 rounded shadow">
          <Form.Group controlId="formContent">
            <Form.Label>Content:</Form.Label>
            <Form.Control as="textarea" name="content" value={formData.content} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formSolution">
            <Form.Label>Solution:</Form.Label>
            <Form.Control as="textarea" name="solution" value={formData.solution} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="formProfile">
            <Form.Label>Profile:</Form.Label>
            <Form.Select name="profile" value={formData.profile} onChange={handleChange}>
              <option value="">Select a profile</option>
              {profiles.map((profile, index) => (
                <option key={index} value={profile.name}>{profile.name}</option>
              ))}
            </Form.Select>
          </Form.Group>
          <Form.Group controlId="formCategory">
            <Form.Label>Category:</Form.Label>
            <Form.Select name="category" value={formData.category} onChange={handleChange}>
              <option value="">Select a category</option>
              {categories.map((category, index) => (
                <option key={index} value={category}>{category}</option>
              ))}
            </Form.Select>
          </Form.Group>
          <Button type="submit">Add Prompt</Button>
        </Form>
        <ListGroup>
          {prompts.map((prompt, index) => (
            <ListGroup.Item key={index} action onClick={() => handlePromptClick(prompt)}>
              {prompt.content}
            </ListGroup.Item>
          ))}
        </ListGroup>
        <DetailPopup item={selectedPrompt} onClose={handleClosePopup} />
      </Container>
    </Layout>
  );
}
