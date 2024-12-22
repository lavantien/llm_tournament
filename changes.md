# Change Log

## Initial Setup
- **Summary**: Created the initial structure for the Next.js application with pages for Model Manager, Profile Manager, Prompt Manager, and Leaderboard.
- **Files Changed**:
  - `pages/index.js`
  - `pages/model-manager.js`
  - `pages/profile-manager.js`
  - `pages/prompt-manager.js`
  - `pages/leaderboard.js`
  - `components/Layout.js`
  - `styles/globals.css`
  - `decision.md`
  - `changes.md`

## Update Layout Component
- **Summary**: Updated the Layout component to include a title and improve navigation styling. Created the SQLite database configuration.
- **Files Changed**:
  - `components/Layout.js`
  - `styles/Layout.module.css`
  - `lib/db.js`
  - `changes.md`

## Implement Pages
- **Summary**: Implemented the pages for Model Manager, Profile Manager, Prompt Manager, and Leaderboard with forms for CRUD operations and a sortable table for the leaderboard.
- **Files Changed**:
  - `pages/model-manager.js`
  - `pages/profile-manager.js`
  - `pages/prompt-manager.js`
  - `pages/leaderboard.js`
  - `changes.md`

## Update Profile Field in Prompt Manager
- **Summary**: Updated the profile field in the Prompt Manager page to be a dropdown selection from added profiles.
- **Files Changed**:
  - `pages/profile-manager.js`
  - `pages/prompt-manager.js`
  - `changes.md`

## Implement Detail Popup
- **Summary**: Implemented the detail popup for all pages to display details when clicking on an item.
- **Files Changed**:
  - `components/DetailPopup.js`
  - `styles/DetailPopup.module.css`
  - `pages/model-manager.js`
  - `pages/profile-manager.js`
  - `pages/prompt-manager.js`
  - `pages/leaderboard.js`
  - `changes.md`
