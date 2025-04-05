# blogapp-microservices

# Functional Requirements
* A user should be able to register and login
* A user should be able to update his/her profile and add an avatar
* A user should be able to create/update/delete blogs
* A user should be able to follow other users
* A user should have a **_home feed_** where the user should see blogs of people he/she follows
* When a blog is created/updated, it should be flagged if there are inappropriate content
* When a user creates a blog, the new blog should appear in followers' feed

# Non-functional Requirements
* Latency (P95):
  * Login: < 100ms
  * Feed load: < 200ms
  * Create blog: < 300ms
* Throughput:
  * 100 req/sec sustained
  * 500 concurrent users peak
* Scale target:
  * 10k active users
  * 500k total blogs
  * 1M follow relationships

# System Architecture

# Tech Stack
* Golang + Echo
* PostgreSQL
* Docker + Kubernetes
* Terraform
* AWS S3
* AWS EKS
* ArgoCD
* Prometheus + Grafana
* Grafaka k6
