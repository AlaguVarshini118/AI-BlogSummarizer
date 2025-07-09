# 🧠 AI Blog Summarizer ✨

A student-friendly web platform that fetches blogs from TechCrunch or lets users paste any blog link to generate concise AI-based summaries. It also includes a powerful vocabulary learning assistant that extracts new words, shows definitions & synonyms, and lets users save words for later revision.

---

## 🚀 Features

- Read blogs directly from TechCrunch
- Paste any blog link to get an AI-generated summary
- Vocabulary section: click on words to view definitions, synonyms, and save them
- User authentication via JWT & Google OAuth2
- AI summarization powered by Cohere API
- Save and manage learned words
- Responsive UI with Tailwind CSS
- Deployed using Render, Docker, and GitHub Actions

---

## 🛠️ Tech Stack

Frontend: React.js, Redux Toolkit, Tailwind CSS, Axios  
Backend: Golang (Gin), PostgreSQL  
AI: Cohere API (for summarization), WordsAPI (for vocabulary)  
Auth: JWT, Google OAuth2  
Deployment: Docker, Render, GitHub Actions  

---

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/AlaguVarshini118/AI-BlogSummarizer.git
cd AI-BlogSummarizer

# Frontend setup
cd frontend
npm install
npm run dev

# Backend setup
cd ../backend
go mod tidy
go run main.go
