## Git Workflow and Issue Management Guidelines

This document outlines the Git workflow and issue management practices for the
project. It is designed to maintain consistency, clarity, and ease of
collaboration, especially as the project evolves.

### **Commit Message Format**

```
feat:<required short-description>
<optional-description>
<optional-reference>
```

#### **Commit Message Components**

1.  **Type**: The type describes what the commit is doing in relation to the
    codebase. Some common commit types include:

    - **`feat`**: A new feature or enhancement to existing functionality. Use
      this when adding a new feature or significant improvement.

    - **`fix`**: A bug fix or a correction of unexpected behavior.

    - **`docs`**: Changes to documentation, including updates to README,
      comments, etc.

    - **`chore`**: Maintenance tasks like updating dependencies, refactoring, or
      other minor improvements.

    - **`style`**: Changes that don’t affect the functionality of the code (e.g.,
      formatting, white-space, missing semicolons).

    - **`refactor`**: A change that modifies the code structure or improves
      readability without changing functionality.

    - **`test`**: Changes related to tests, such as adding or modifying test
      cases.

2.  **Short Description**:

    - The description should be **concise** and written in the **imperative
      mood** (i.e., "Add", "Fix", "Remove" instead of "Added", "Fixes", "Removed").

    - Keep the description to around 50-72 characters, ensuring it’s brief but
      informative enough to convey the essence of the change.

      ```markdown
      feat: add profile picture upload feature
      ```

3.  **Detailed Description** (Optional but Recommended for Complex Changes):

    - If your commit involves complex logic or multiple steps, provide additional
      context. This helps anyone reviewing your commit understand the "why" behind
      the change, and it's especially useful if the change might not be immediately
      obvious.
    - Keep the detailed description **short and to the point**. Focus on what was
      done and why, rather than how.

      ```markdown
      feat: add user profile page

      This commit introduces a new user profile page where users can update their
      personal information, change their password, and upload a profile picture.
      ```

4.  **Issue Reference (Optional but Recommended)**:

    - If your commit addresses an issue or task, reference the issue in the
      message. Use GitHub’s auto-closing keywords like `Fixes`, `Resolves`, or
      `Closes` followed by the issue number to automatically close the issue when
      the commit is merged.

      ```markdown
      feat: add user profile page

      Fixes: #123<auto-closing>
      ```

    - If the commit is part of a larger feature but doesn't resolve the issue
      directly, you can reference the issue as a reference:

      ```markdown
      feat: add user profile page

      Reference: #123<auto-closing>
      ```

---

#### **Best Practices for Commit Messages**

- **Be concise but descriptive**: The commit message should briefly explain what
  was done, but enough to understand the essence of the change.

- **Use the appropriate type**: This helps with organizing commits and gives
  context to your team or collaborators.

- **Refer to issues or tasks**: Always link commits to issues where possible.
  This creates a clear link between the code changes and the problem they solve.

- **Imperative mood** means writing your commit messages as commands. This style
  is consistent with how Git treats commit messages and helps to align with how
  the history is interpreted.This ensures that commit messages read like a
  command, as if you’re telling Git
  what to do, making them more straightforward and easier to follow in the context
  of a project's history.

  - **Good**: `fix: correct typo in registration form`
  - **Bad**: `fixed: corrected typo in registration form`

---

### **Issue Format**

1.  **Prefix Categories**: Use prefixes to categorize issues based on their
    type. This provides an immediate understanding of the issue's focus. Common
    prefixes include:

    - **[Bug]**: Use this prefix for issues related to bugs, errors, or unexpected
      behaviors.

    - **[Feature]**: Use this prefix for new feature requests or enhancements.

    - **[Task]**: Use this prefix for general tasks, such as refactoring, updates, or documentation improvements.

    - **[Docs]**: Use this prefix for issues related to documentation.

    - **[Chore]**: Use this prefix for maintenance or routine tasks.

2.  **Short Description**: The issue title should provide a brief but clear
    description of the problem or feature. Aim for clarity and keep it
    concise—ideally, under 72 characters.

    - **Bug Example**: `[Bug] Fix issue with form validation on login page`
    - **Feature Example**: `[Feature] Add dark mode toggle to the UI`
    - **Task Example**: `[Task] Refactor authentication logic`

3.  **Use of Action Verbs**: Use action verbs in the issue title to describe the
    problem or desired outcome. This helps set expectations about what needs to be
    done.

    - **Example**: `Fix`, `Add`, `Improve`, `Update`, `Refactor`, `Resolve`,
      `Remove`, etc.
    - **Example**: `[Bug] Fix crash when submitting form without required fields`

4.  **Reference Problem or Context**: When possible, include the specific
    problem, context, or feature being worked on. This makes the issue easier to
    understand at a glance.

    - **Bug Example**: `[Bug] Fix incorrect error message on invalid email input`
    - **Feature Example**: `[Feature] Add email notifications for user registration`

5.  **Prioritization and Labels**: If your issues are large, break them down
    into smaller tasks or bugs, and label the issue with priority tags like `high`,
    `medium`, or `low`. You can also use labels like `needs review`, `in progress`,
    or `blocked` to keep track of the status.

---

**Example Workflow in Practice**

1.  **Bug**:

    - Commit message: `fix: correct email validation logic`
    - Issue: `[Bug] Fix issue with incorrect email validation`

2.  **Feature**:

    - Commit message: `feat: add search bar to homepage `
    - Issue: `[Feature] Add search bar to homepage`
