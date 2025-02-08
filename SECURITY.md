# Security Policy

Thank you for helping to keep my project secure! This document provides guidelines on **responsible disclosure**, **security best practices**, and **how to report vulnerabilities**.

---

## **Supported Versions**
I actively maintain the latest version of the project. Below is my support policy:

| Version  | Supported |
|----------|-----------|
| `main` (latest) | ✅ Yes |
| Older releases | ❌ No |

Security patches are applied **only to the latest version**. Please ensure you're using an up-to-date release.

---

## **Reporting a Vulnerability**
If you discover a security issue, **please report it responsibly**:

- **Contact:** [nassyrovich@gmail.com](mailto:nassyrovich@gmail.com)  
- **DO NOT** disclose vulnerabilities publicly until I have released a fix.  
- I aim to acknowledge reports within **48 hours** and release patches within **7 days**, depending on severity.  

If needed, I can provide **CVE assignment** for high-impact vulnerabilities.

---

## **Security Best Practices**
To ensure the security of my project, I follow these best practices:

### **1. Secure Development**
- **Do not commit secrets** (API keys, credentials). Use `.gitignore` to exclude them.  
- **Enable 2FA** (two-factor authentication) on GitHub/GitLab accounts.  
- **Use signed commits** (`git commit -S`) to ensure authenticity.  
- **Keep dependencies updated** (Dependabot, Renovate).  

### **2. Secure CI/CD Pipelines**
- Store sensitive credentials in **GitHub Secrets / GitLab Variables**.  
- **Use least privilege access** for CI/CD runners.  
- Scan code & dependencies with **SAST (CodeQL, SonarQube, Semgrep)**.  

### **3. Secure Deployment**
- Deploy using **Docker & Kubernetes with least privilege roles**.  
- Ensure **HTTPS everywhere** (TLS certificates via Let's Encrypt).  
- Scan Docker images (`trivy`, `grype`) before deployment.  

---

## **Responsible Disclosure Process**
1. **Submit a report** privately via email.  
2. I will **confirm receipt** and investigate the issue.  
3. If valid, I will **assign a severity level** and plan a fix.  
4. A **patch will be released**, and you will be credited (if desired).  
5. After release, I may **publicly disclose the vulnerability**.  

For security-related questions, feel free to **contact me anytime!**

---

## **License & Disclaimer**
This security policy follows **best practices** and is subject to change as the project evolves.
