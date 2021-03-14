describe('Mailing forms', () => {
  it('checks labels and validation of a form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // Ensure the mailing form exists in the Engage page
    cy.get('[data-cy=mailing-form]');
    // By referencing general-tab items we are ensuring that it is selected by default
    // Check if name label exists and is required
    cy.get('[data-cy=general-name-label]').contains('Name').should('have.class', 'required');
    // Check if email label exists and is required
    cy.get('[data-cy=general-email-label]').contains('Email').should('have.class', 'required');
    // Check if message label exists and is required
    cy.get('[data-cy=general-message-label]').contains('Message').should('have.class', 'required');
    // Ensure send button is disabled
    cy.get('[data-cy=general-send-button]').should('be.disabled');
    // Fill in valid name
    cy.get('[data-cy=general-name-field]').type('Sergio');
    // Fill in invalid email
    cy.get('[data-cy=general-email-field]').type('sergio2@ema.il.c');
    // Fill in valid message
    cy.get('[data-cy=general-message-field]').type('message goes here');
    // Ensure send button is disabled,
    // because the email is invalid
    cy.get('[data-cy=general-send-button]').should('be.disabled');
    // Fill in valid email
    cy.get('[data-cy=general-email-field]').clear().type('sergio@gmail.com');
    // Ensure send button is enabled after making the email valid
    cy.get('[data-cy=general-send-button]').should('not.be.disabled');
  });

  it('ensures the name label is different in a sponsorship form', () => {
    // Visit sponsors page
    cy.visit('/#/sponsors');
    // Ensure the mailing form exists in the Sponsors page
    cy.get('[data-cy=mailing-form]');
    // Check if name label changed to the sponsorship form version
    cy.get('[data-cy=sponsorship-name-label]').contains('Company Name').should('have.class', 'required');
  });

  it('checks labels and validation of a feedback form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // Select feedback form tab
    cy.get('[data-cy=feedback-form-selector]').click();
    // Check if name label is not required (marked with *)
    cy.get('[data-cy=feedback-name-label]').should('not.have.class', 'required');
    /// Check if email label is not required (marked with *)
    cy.get('[data-cy=feedback-email-label]').should('not.have.class', 'required');
    // Ensure send button is disabled
    cy.get('[data-cy=feedback-send-button]').should('be.disabled');
    // Fill in valid message
    cy.get('[data-cy=feedback-message-field]').type('message goes here');
    // Ensure send button is enabled,
    // because message is the only required field
    cy.get('[data-cy=feedback-send-button]').should('not.be.disabled');
  });
});
