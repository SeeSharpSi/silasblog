// Enhanced JavaScript for Silas Tompkins Blog
// Handles animations, interactions, and enhanced UX

document.addEventListener("DOMContentLoaded", function () {
  // Initialize all functionality
  initializeAnimations();
  initializeIntersectionObserver();
  initializeScrollEffects();
  initializeSmoothScrolling();
  initializeInteractiveElements();
});

// Handle HTMX after request events
document.addEventListener("htmx:afterRequest", function () {
  // Re-initialize animations after HTMX requests
  initializeAnimations();
  initializeIntersectionObserver();

  // Smooth scroll to top after page transitions
  window.scrollTo({ top: 0, behavior: "smooth" });
});

// Initialize hero section animations
function initializeAnimations() {
  // Hero stats counter animation
  const stats = document.querySelectorAll(".stat-number");
  stats.forEach((stat) => {
    const text = stat.textContent;
    const number = parseInt(text.replace(/\D/g, ""));

    if (number && !stat.classList.contains("animated")) {
      stat.classList.add("animated");
      animateCounter(stat, 0, number, text);
    }
  });

  // Hero text stagger animation
  const heroElements = document.querySelectorAll(
    ".hero-title, .hero-subtitle, .hero-description",
  );
  heroElements.forEach((element, index) => {
    element.style.opacity = "0";
    element.style.transform = "translateY(30px)";

    setTimeout(() => {
      element.style.transition = "all 0.8s cubic-bezier(0.4, 0, 0.2, 1)";
      element.style.opacity = "1";
      element.style.transform = "translateY(0)";
    }, index * 200);
  });

  // Legacy headshot text animation (keep for compatibility)
  const headshotItems = document.querySelectorAll(".headshot_txt div div");
  headshotItems.forEach((div, i) => {
    div.style.opacity = "0";
    div.style.transform = "translateY(20px)";

    setTimeout(() => {
      div.style.transition = "all 0.6s cubic-bezier(0.4, 0, 0.2, 1)";
      div.style.opacity = "1";
      div.style.transform = "translateY(0)";
    }, i * 100);
  });
}

// Counter animation for statistics
function animateCounter(element, start, end, originalText) {
  const duration = 2000;
  const startTime = performance.now();

  function update(currentTime) {
    const elapsed = currentTime - startTime;
    const progress = Math.min(elapsed / duration, 1);

    // Easing function for smooth animation
    const easeOut = 1 - Math.pow(1 - progress, 3);
    const current = Math.floor(start + (end - start) * easeOut);

    // Update the text while preserving formatting
    if (originalText.includes("+")) {
      element.textContent = current + "+";
    } else if (originalText.includes("s")) {
      element.textContent = current + "s";
    } else {
      element.textContent = current.toString();
    }

    if (progress < 1) {
      requestAnimationFrame(update);
    } else {
      element.textContent = originalText; // Ensure final text is correct
    }
  }

  requestAnimationFrame(update);
}

// Intersection Observer for scroll-triggered animations
function initializeIntersectionObserver() {
  const observerOptions = {
    threshold: 0.1,
    rootMargin: "0px 0px -50px 0px",
  };

  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add("animate-in");

        // Special handling for skill tags
        if (entry.target.classList.contains("skill-category")) {
          const skillTags = entry.target.querySelectorAll(".skill-tag");
          skillTags.forEach((tag, index) => {
            setTimeout(() => {
              tag.style.opacity = "1";
              tag.style.transform = "translateY(0) scale(1)";
            }, index * 100);
          });
        }

        // Special handling for experience highlights
        if (entry.target.classList.contains("experience-highlights")) {
          const highlights = entry.target.querySelectorAll("li");
          highlights.forEach((highlight, index) => {
            setTimeout(() => {
              highlight.style.opacity = "1";
              highlight.style.transform = "translateX(0)";
            }, index * 150);
          });
        }
      }
    });
  }, observerOptions);

  // Observe elements for animation
  const elementsToObserve = document.querySelectorAll(`
        .skill-category,
        .experience-card,
        .education-card,
        .contact-link,
        .article_thumbnail,
        .experience-highlights
    `);

  elementsToObserve.forEach((element) => {
    observer.observe(element);

    // Set initial states
    if (element.classList.contains("skill-category")) {
      const skillTags = element.querySelectorAll(".skill-tag");
      skillTags.forEach((tag) => {
        tag.style.opacity = "0";
        tag.style.transform = "translateY(20px) scale(0.9)";
        tag.style.transition = "all 0.4s cubic-bezier(0.4, 0, 0.2, 1)";
      });
    }

    if (element.classList.contains("experience-highlights")) {
      const highlights = element.querySelectorAll("li");
      highlights.forEach((highlight) => {
        highlight.style.opacity = "0";
        highlight.style.transform = "translateX(-20px)";
        highlight.style.transition = "all 0.5s cubic-bezier(0.4, 0, 0.2, 1)";
      });
    }
  });
}

// Parallax and scroll effects
function initializeScrollEffects() {
  let ticking = false;

  function updateScrollEffects() {
    const scrolled = window.pageYOffset;
    const rate = scrolled * -0.5;

    // Parallax effect for hero section
    const heroImage = document.querySelector(".hero-image");
    if (heroImage) {
      heroImage.style.transform = `translateY(${rate * 0.1}px)`;
    }

    // Navigation background opacity
    const nav = document.querySelector(".topNav");
    if (nav) {
      const opacity = Math.min(scrolled / 100, 1);
      nav.style.backgroundColor = `rgba(29, 115, 115, ${0.9 + opacity * 0.1})`;
    }

    ticking = false;
  }

  function requestTick() {
    if (!ticking) {
      requestAnimationFrame(updateScrollEffects);
      ticking = true;
    }
  }

  window.addEventListener("scroll", requestTick, { passive: true });
}

// Smooth scrolling for anchor links
function initializeSmoothScrolling() {
  document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
    anchor.addEventListener("click", function (e) {
      e.preventDefault();
      const target = document.querySelector(this.getAttribute("href"));
      if (target) {
        target.scrollIntoView({
          behavior: "smooth",
          block: "start",
        });
      }
    });
  });
}

// Interactive elements and enhanced UX
function initializeInteractiveElements() {
  // Enhanced hover effects for contact links
  const contactLinks = document.querySelectorAll(".contact-link");
  contactLinks.forEach((link) => {
    link.addEventListener("mouseenter", function () {
      this.style.transform = "translateY(-6px) scale(1.02)";
    });

    link.addEventListener("mouseleave", function () {
      this.style.transform = "translateY(0) scale(1)";
    });
  });

  // Skill tag interactions
  const skillTags = document.querySelectorAll(".skill-tag");
  skillTags.forEach((tag) => {
    tag.addEventListener("mouseenter", function () {
      // Add ripple effect
      const ripple = document.createElement("span");
      ripple.className = "ripple";
      this.appendChild(ripple);

      setTimeout(() => {
        ripple.remove();
      }, 600);
    });
  });

  // Article thumbnail hover effects
  const articleThumbnails = document.querySelectorAll(".article_thumbnail");
  articleThumbnails.forEach((thumbnail) => {
    thumbnail.addEventListener("mouseenter", function () {
      this.style.transform = "translateY(-8px) scale(1.02)";
    });

    thumbnail.addEventListener("mouseleave", function () {
      this.style.transform = "translateY(0) scale(1)";
    });
  });

  // Headshot image interaction
  const headshot = document.querySelector(".headshot");
  if (headshot) {
    headshot.addEventListener("mouseenter", function () {
      this.style.transform = "scale(1.1) rotate(2deg)";
    });

    headshot.addEventListener("mouseleave", function () {
      this.style.transform = "scale(1) rotate(0deg)";
    });
  }
}

// Utility function to add CSS animations
function addAnimationStyles() {
  const style = document.createElement("style");
  style.textContent = `
        .animate-in {
            animation: slideInUp 0.8s cubic-bezier(0.4, 0, 0.2, 1) forwards;
        }

        @keyframes slideInUp {
            from {
                opacity: 0;
                transform: translateY(30px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }

        .ripple {
            position: absolute;
            border-radius: 50%;
            background: rgba(255, 255, 255, 0.4);
            transform: scale(0);
            animation: rippleEffect 0.6s linear;
            pointer-events: none;
        }

        @keyframes rippleEffect {
            to {
                transform: scale(4);
                opacity: 0;
            }
        }
    `;
  document.head.appendChild(style);
}

// Initialize CSS animations
addAnimationStyles();

// Performance optimization: Debounce function
function debounce(func, wait) {
  let timeout;
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout);
      func(...args);
    };
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
  };
}

// Handle window resize events
window.addEventListener(
  "resize",
  debounce(() => {
    // Recalculate animations on resize
    initializeAnimations();
  }, 250),
);

// Add loading states for better UX
document.addEventListener("htmx:beforeRequest", function () {
  document.body.classList.add("loading");
});

document.addEventListener("htmx:afterRequest", function () {
  document.body.classList.remove("loading");
});

// Accessibility improvements
document.addEventListener("keydown", function (e) {
  // Allow tabbing through interactive elements
  if (e.key === "Tab") {
    document.body.classList.add("keyboard-nav");
  }
});

document.addEventListener("mousedown", function () {
  document.body.classList.remove("keyboard-nav");
});

// Console message for developers
console.log("🚀 Silas Tompkins Blog - Enhanced with modern JavaScript");
console.log("💻 Built with Go, HTMX, and Templ");
console.log("🎨 Styled with modern CSS and smooth animations");
