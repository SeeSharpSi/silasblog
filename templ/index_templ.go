// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.906
package templ

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><title>Silas Tompkins - Software Engineer</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><link rel=\"stylesheet\" type=\"text/css\" href=\"/static/styles.css\"><script type=\"text/javascript\" src=\"/static/htmx.min.js\"></script><script src=\"/static/script.js\"></script><style>\n        body {\n            font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;\n            margin: 0;\n            padding: 0;\n        }\n    </style></head><body><div class=\"topNav\" hx-boost=\"true\"><a href=\"/\" hx-select=\"#page\" hx-target=\"#page\">Home</a> <a href=\"/articles\" hx-select=\"#page\" hx-target=\"#page\">Articles</a></div><div id=\"page\"><!-- Hero Section --><section class=\"hero\"><div class=\"hero-content\"><div class=\"hero-text\"><h1 class=\"hero-title\">Silas Tompkins</h1><h2 class=\"hero-subtitle\">Software Engineer</h2><p class=\"hero-description\">Software Engineering student with over two years of hands-on experience in enterprise application development, CI/CD pipelines, and cloud infrastructure. Developed systems used by thousands of users.</p><div class=\"hero-stats\"><div class=\"stat\"><span class=\"stat-number\">3,000+</span> <span class=\"stat-label\">Users Served</span></div><div class=\"stat\"><span class=\"stat-number\">2+</span> <span class=\"stat-label\">Years Experience</span></div><div class=\"stat\"><span class=\"stat-number\">100s</span> <span class=\"stat-label\">Repos Migrated</span></div></div></div><div class=\"hero-image\"><img class=\"headshot\" src=\"/static/headshot.jpg\" alt=\"Silas Tompkins\"></div></div></section><!-- Skills Section --><section class=\"skills-section\"><div class=\"container\"><h2 class=\"section-title\">Technical Skills</h2><div class=\"skills-grid\"><div class=\"skill-category\"><h3>Languages</h3><div class=\"skill-tags\"><span class=\"skill-tag\">Go</span> <span class=\"skill-tag\">Python</span> <span class=\"skill-tag\">JavaScript</span> <span class=\"skill-tag\">TypeScript</span> <span class=\"skill-tag\">Java</span> <span class=\"skill-tag\">PHP</span></div></div><div class=\"skill-category\"><h3>Frameworks & Tools</h3><div class=\"skill-tags\"><span class=\"skill-tag\">Vue.js</span> <span class=\"skill-tag\">Laravel</span> <span class=\"skill-tag\">Spring Boot</span> <span class=\"skill-tag\">Docker</span> <span class=\"skill-tag\">CI/CD</span> <span class=\"skill-tag\">SQL</span></div></div><div class=\"skill-category\"><h3>Cloud & DevOps</h3><div class=\"skill-tags\"><span class=\"skill-tag\">AWS</span> <span class=\"skill-tag\">GitLab</span> <span class=\"skill-tag\">GitHub</span> <span class=\"skill-tag\">Linux</span> <span class=\"skill-tag\">Bash</span></div></div></div></div></section><!-- Experience Section --><section class=\"experience-section\"><div class=\"container\"><h2 class=\"section-title\">Recent Experience</h2><div class=\"experience-card\"><div class=\"experience-header\"><h3>Enterprise Applications Co-op</h3><span class=\"company\">Georgia Tech Research Institute</span> <span class=\"duration\">May 2023 – May 2025</span></div><ul class=\"experience-highlights\"><li>Developed a full-stack web application used by over <strong>3,000 employees</strong><br><span class=\"duration\">PHP, Javascript, CI/CD, Docker</span></li><li>Automated migration of <strong>hundreds of repositories</strong> from BitBucket to GitLab<br><span class=\"duration\">Go, Python</span></li><li>Built security tools to scan Git repositories for exposed secrets<br><span class=\"duration\">Python</span></li><li>Maintained uptime for Atlassian Confluence, Jira, and Service Management<br><span class=\"duration\">Linux, Bash</span></li><li>Gained experience with NIST 800-171 compliant systems</li></ul></div></div><div class=\"container\"><div class=\"experience-card\" style=\"margin-top: 16px\"><div class=\"experience-header\"><h3>Web Design Intern</h3><span class=\"company\">Digital Explorations</span> <span class=\"duration\">February 2023 – April 2023</span></div><ul class=\"experience-highlights\"><li>Created user interfaces with Vue.js and JavaScript<br><span class=\"duration\">Vue.js, JavaScript, TypeScript</span></li><li>Collaborated with a team using Figma and GitHub<br><span class=\"duration\">CI/CD, GitHub</span></li><li>Connected back-end and front-end systems using AWS Amplify and NoSQL<br><span class=\"duration\">AWS, SQL</span></li></ul></div></div></section><!-- Education Section --><section class=\"education-section\"><div class=\"container\"><h2 class=\"section-title\">Education</h2><div class=\"education-card\"><h3>Bachelor of Science in Software Engineering</h3><div class=\"education-details\"><span class=\"university\">Kennesaw State University</span> <span class=\"gpa\">GPA: 3.44</span> <span class=\"graduation\">Graduated: May 2025</span></div></div></div></section><!-- Contact Section --><section class=\"contact-section\"><div class=\"container\"><h2 class=\"section-title\">Get In Touch</h2><div class=\"contact-links\"><a href=\"mailto:silastompkins@outlook.com\" class=\"contact-link\"><span class=\"contact-icon\">✉</span> <span>silastompkins@outlook.com</span></a> <a href=\"https://www.linkedin.com/in/silas-tompkins\" class=\"contact-link\" target=\"_blank\"><span class=\"contact-icon\">💼</span> <span>LinkedIn</span></a> <a href=\"https://github.com/SeeSharpSi\" class=\"contact-link\" target=\"_blank\"><span class=\"contact-icon\">💻</span> <span>GitHub</span></a></div></div></section><!-- Status Section --><section class=\"status-section\"><div class=\"container\"><div class=\"status-card\"><p>🚧 This website is currently under development. Track progress at <a href=\"https://github.com/seesharpsi/silasblog\" target=\"_blank\">github.com/seesharpsi/silasblog</a></p></div></div></section></div><div id=\"cookies\">This site doesn't use cookies.<br>(but really, most sites shouldn't) <span class=\"close_cookies_btn\" onclick=\"this.parentElement.style.display='none'\">&times;</span></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
