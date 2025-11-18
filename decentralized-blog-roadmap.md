# Decentralized Blogging Platform - Student Project Roadmap

## Project Overview
A resume-worthy project showcasing modern web development with Go and HTMX. Build a personal blogging platform that can connect with friends' instances to share posts.

**Why This Project Stands Out:**
- Shows full-stack development skills (Go backend, HTMX frontend)
- Demonstrates understanding of distributed systems
- Great talking point in interviews
- Actually useful - you and friends can use it!

---

## 🎯 MVP - Minimum Viable Product (2-3 weeks)

Start here! This is enough for a solid resume project.

### Week 1: Basic Blog (Core Features)
**What you're building:** A simple personal blog that works locally.

- [x] ~~Set up Go project with basic structure~~
- [x] ~~Create SQLite database (super simple, no install needed)~~
- [ ] Make a homepage that lists blog posts
- [ ] Build "Create Post" page with a form
- [ ] Build "View Post" page to read a single post
- [ ] Add markdown support (use a Go library like `goldmark`)
- [ ] Style with basic CSS (can use Tailwind CDN for quick styling)

**Resume Bullet:** "Built full-stack blogging application using Go and SQLite with server-side rendering"

### Week 2: HTMX Interactivity
**What you're building:** Make it feel modern without JavaScript frameworks.

- [ ] Add HTMX to your HTML templates
- [ ] Make post creation work without page refresh
- [ ] Add edit/delete buttons (that work with HTMX)
- [ ] Add simple search that filters posts (with HTMX)
- [ ] Show "Loading..." states when things are happening

**Resume Bullet:** "Implemented dynamic UI interactions using HTMX for seamless user experience"

### Week 3: Make It Look Good & Deploy
**What you're building:** Polish for your resume/portfolio.

- [ ] Add profile page with bio and profile picture
- [ ] Make it responsive (works on phone/tablet)
- [ ] Add dark mode toggle
- [ ] Write a good README with screenshots
- [ ] Deploy to a free service (Fly.io, Railway, or Render)
- [ ] Add Docker support (optional but looks great)

**Resume Bullet:** "Deployed containerized web application with Docker to cloud platform"

---

## 🚀 Phase 2 - Federation (Optional, 2-3 weeks)

Add this if you have time and want to really impress!

### Week 4-5: Basic Federation
**What you're building:** Connect your blog to a friend's blog.

- [ ] Create an RSS feed of your posts (super simple)
- [ ] Add "Following" feature to follow other blogs
- [ ] Fetch and display posts from blogs you follow
- [ ] Create a combined feed showing your posts + followed posts
- [ ] Add simple authentication (just username/password)

**Resume Bullet:** "Implemented RSS-based federation protocol enabling cross-instance content sharing"

### Week 6: Polish Federation
**What you're building:** Make the federation actually work well.

- [ ] Add "Follow" button that works with any blog URL
- [ ] Show avatars/names from other instances
- [ ] Add comment system (comments stored locally)
- [ ] Make it easy to reply to posts from other instances
- [ ] Update documentation with federation setup

**Resume Bullet:** "Designed and implemented decentralized social features for multi-instance communication"

---

## 📚 Tech Stack (Keep It Simple!)

**Backend:**
- Go (standard library is mostly enough!)
- SQLite (comes with Go, super easy)
- `html/template` (built into Go)

**Frontend:**
- HTMX (one `<script>` tag!)
- Tailwind CSS (via CDN, or just regular CSS)
- No React, No Vue, No complex build process

**Deployment:**
- Docker (write a simple Dockerfile)
- Free hosting: Fly.io, Railway, or Render
- GitHub for version control

---

## 🎓 Learning Goals (Great for Interviews!)

This project teaches you:

1. **Backend Development** - Go, HTTP servers, databases
2. **Frontend Without Framework** - HTMX, progressive enhancement
3. **Database Design** - Schema design, queries, indexes
4. **Deployment** - Docker, cloud platforms, environment configs
5. **Distributed Systems** - Federation, data synchronization (if you do Phase 2)

---

## 💡 Interview Talking Points

**What I Built:**
"A decentralized blogging platform where users can host their own instance and follow each other. Built with Go for the backend and HTMX for interactive UI without heavy JavaScript frameworks."

**Technical Challenges:**
- "Designed a simple federation protocol using RSS feeds"
- "Optimized database queries for listing and searching posts"
- "Implemented server-side rendering with partial updates using HTMX"

**What I Learned:**
- "How to build full-stack applications without JavaScript frameworks"
- "Database schema design and optimization"
- "Containerization and cloud deployment"
- "Basic distributed systems concepts"

---

## ✅ Week-by-Week Checklist

### Week 1: Foundation
- [ ] Day 1-2: Setup Go, create basic HTTP server, design database
- [ ] Day 3-4: Build post creation and listing pages
- [ ] Day 5-7: Add markdown rendering, basic styling

### Week 2: HTMX
- [ ] Day 1-2: Add HTMX, make forms work without page reload
- [ ] Day 3-4: Add edit/delete with HTMX
- [ ] Day 5-7: Add search and loading states

### Week 3: Polish
- [ ] Day 1-2: Profile page, make it responsive
- [ ] Day 3-4: Add dark mode, improve styling
- [ ] Day 5-7: Write README, deploy, take screenshots

### Week 4-5: Federation (Optional)
- [ ] Days 1-5: RSS feed, following system
- [ ] Days 6-10: Combined feed, basic auth

### Week 6: Final Polish (Optional)
- [ ] Days 1-3: Comments, reply system
- [ ] Days 4-7: Documentation, demo video

---

## 🎬 Demo Ideas for Your Portfolio

1. **Live Demo:** Host your instance, share the URL
2. **Video Demo:** 2-3 minute walkthrough showing:
   - Creating a post
   - HTMX in action (no page reloads)
   - Following another instance
   - Viewing combined feed
3. **GitHub README:** Include:
   - Screenshots
   - Architecture diagram (simple boxes and arrows)
   - Quick start guide
   - What you learned section

---

## 🔧 Minimum Code Structure

```
decentralized-blog/
├── main.go              # HTTP server setup
├── handlers.go          # Route handlers
├── database.go          # Database queries
├── models.go            # Post, User structs
├── templates/           # HTML templates
│   ├── layout.html
│   ├── home.html
│   ├── post.html
│   └── create.html
├── static/              # CSS, images
├── README.md
└── Dockerfile
```

Keep it simple! You can expand later.

---

## 🎯 Success Criteria

**Minimum (for resume):**
- ✅ Works on your computer
- ✅ Can create, edit, delete posts
- ✅ Uses HTMX for some interactions
- ✅ Deployed somewhere with a live URL
- ✅ Clean code on GitHub with good README

**Impressive (if you have time):**
- ✅ Federation works with 2+ instances
- ✅ Good documentation
- ✅ Docker setup
- ✅ Responsive design
- ✅ Demo video

---

## 💪 Stay Motivated!

**Week 1:** "I built a working blog!"
**Week 2:** "It's interactive without React!"
**Week 3:** "It's deployed and looks professional!"
**Week 4-6:** "Multiple instances can talk to each other!"

Remember: Even just the MVP (weeks 1-3) is impressive for a student project. Don't burn out trying to do everything!

---

## 🤔 Keep It Simple Guidelines

**Do:**
- Use SQLite (not PostgreSQL)
- Use Go standard library (minimal dependencies)
- Use HTMX from CDN (no npm/webpack)
- Start with single-user instances
- Deploy to free tier services

**Don't (at first):**
- Try to implement ActivityPub (too complex)
- Build a mobile app
- Add user authentication (single-user is fine)
- Worry about scaling to thousands of users
- Over-engineer the federation protocol

---

## 📅 Timeline Summary

- **3 weeks = Good resume project**
- **6 weeks = Excellent portfolio project**
- **8+ weeks = Talking point for every interview**

Start with 3 weeks, see how it goes!

---

**Good luck! You've got this! 🚀**

*Remember: The best project is one you finish. Start simple, make it work, then add cool features.*
