# Recommended Tech Stack for LLM Benchmarking App

## Core Technologies

### Frontend & Desktop Framework

- **Tauri + React**
  - Tauri provides native performance with tiny bundle sizes
  - React for robust UI development
  - Better performance than Electron
  - Smaller bundle size and memory footprint
  - Built-in security features

### UI Component Library

- **shadcn/ui**
  - Modern, clean aesthetic
  - Dark theme support out of the box
  - Highly customizable components
  - Built on top of Tailwind CSS
  - Perfect for desktop applications

### State Management

- **Zustand**
  - Lightweight and simple
  - Perfect for reactive updates
  - Easy integration with React
  - Better performance than Redux for this scale

### Database

- **SQLite via better-sqlite3**
  - Perfect for local desktop apps
  - No server required
  - ACID compliant
  - Fast performance
  - Simple setup and maintenance

### Charts and Visualization

- **Recharts**
  - React-based
  - Smooth animations
  - Responsive design
  - Good performance with large datasets
  - Customizable themes

## Project Structure

```
src/
├── components/           # Reusable UI components
├── db/                  # Database operations
│   ├── schema.ts       # Database schema definitions
│   └── operations.ts   # CRUD operations
├── features/           # Feature-specific components
│   ├── home/
│   ├── models/
│   ├── profiles/
│   ├── prompts/
│   ├── leaderboard/
│   └── stats/
├── store/              # Zustand store
├── types/              # TypeScript definitions
└── utils/              # Utility functions
```

## Implementation Details

### Data Layer

```typescript
// Example schema using better-sqlite3
const schema = {
  bots: `
    CREATE TABLE IF NOT EXISTS bots (
      name TEXT PRIMARY KEY,
      path TEXT,
      size REAL,
      param REAL,
      quant TEXT,
      gpuLayers INTEGER,
      gpuLayersUsed INTEGER,
      ctx INTEGER,
      ctxUsed INTEGER,
      kingOf TEXT,
      FOREIGN KEY(kingOf) REFERENCES profiles(name)
    )
  `,
  // ... other tables
};
```

### State Management

```typescript
// Example Zustand store
import create from "zustand";

interface AppState {
  bots: Bot[];
  profiles: Profile[];
  scores: Score[];
  fetchBots: () => Promise<void>;
  updateScores: (scores: Score[]) => Promise<void>;
  // ... other actions
}

const useStore = create<AppState>((set) => ({
  bots: [],
  profiles: [],
  scores: [],
  fetchBots: async () => {
    const bots = await window.__TAURI__.invoke("get_bots");
    set({ bots });
  },
  // ... other implementations
}));
```

## Development Tools

### Required

- Node.js & npm/yarn
- Rust (for Tauri)
- TypeScript
- VS Code with recommended extensions:
  - Rust Analyzer
  - Tauri
  - ESLint
  - Prettier

### Optional but Recommended

- SQLite Browser for database inspection
- React Developer Tools
- Performance monitoring tools

## Performance Considerations

1. Use virtual scrolling for long lists
2. Implement proper database indexing
3. Batch database operations
4. Use proper memoization for expensive calculations
5. Lazy load components when possible

## Testing Strategy

1. Unit tests with Vitest
2. Integration tests with Testing Library
3. E2E tests with Playwright
4. Database operations testing

## Development Workflow

1. Use conventional commits
2. Implement proper error handling
3. Add logging using Tauri's logging facility
4. Use TypeScript strict mode
5. Implement proper data validation
