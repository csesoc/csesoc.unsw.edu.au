describe('Mailing forms', () => {
  it('checks labels and validation of a general form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // By referencing general-tab items we are ensuring that it is selected by default
    // Check if name label exists
    cy.get('[data-cy=general-name-label]').contains('Name *');
    // Check if email label exists
    cy.get('[data-cy=general-email-label]').contains('Email *');
    // Check if message label exists
    cy.get('[data-cy=general-message-label]').contains('Message *');
    // Ensure send button is disabled
    cy.get('[data-cy=general-send-button]').should('be.disabled');
    // Fill in valid name
    cy.get('[data-cy=general-name-field]').type('Sergio');
    // Fill in valid email
    cy.get('[data-cy=general-email-field]').type('sergio@gmail.com');
    // Fill in valid message
    cy.get('[data-cy=general-message-field]').type('message goes here');
    // Ensure send button is enabled after entering valid inputs
    cy.get('[data-cy=general-send-button]').should('not.be.disabled');
  });

  it('checks labels and validation of a sponsorship form', () => {
    // Visit sponsors page
    cy.visit('/#/sponsors');
    // Check if name label exists
    cy.get('[data-cy=sponsorship-name-label]').contains('Company Name *');
    // Check if email label exists
    cy.get('[data-cy=sponsorship-email-label]').contains('Email *');
    // Check if message label exists
    cy.get('[data-cy=sponsorship-message-label]').contains('Message *');
    // Ensure send button is disabled
    cy.get('[data-cy=sponsorship-send-button]').should('be.disabled');
    // Fill in valid name
    cy.get('[data-cy=sponsorship-name-field]').type('Sergio');
    // Fill in invalid email
    cy.get('[data-cy=sponsorship-email-field]').type('sergio2@ema.il.c');
    // Fill in valid message
    cy.get('[data-cy=sponsorship-message-field]').type('message goes here');
    // Ensure send button is disabled,
    // because the email is invalid
    cy.get('[data-cy=sponsorship-send-button]').should('be.disabled');
    // Fill in valid email
    cy.get('[data-cy=sponsorship-email-field]').clear().type('sergio@gmail.com');
    // Ensure send button is enabled after making the email valid
    cy.get('[data-cy=sponsorship-send-button]').should('not.be.disabled');
  });

  it('checks labels and validation of a feedback form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // Select feedback form tab
    cy.get('[data-cy=feedback-form-tab]').click();
    // Check if name label exists
    cy.get('[data-cy=feedback-name-label]').contains('Name');
    // Check if email label exists
    cy.get('[data-cy=feedback-email-label]').contains('Email');
    // Check if message label exists
    cy.get('[data-cy=feedback-message-label]').contains('Message *');
    // Ensure send button is disabled
    cy.get('[data-cy=feedback-send-button]').should('be.disabled');
    // Fill in valid message
    cy.get('[data-cy=feedback-message-field]').type('message goes here');
    // Ensure send button is enabled,
    // because the message is the only required field
    cy.get('[data-cy=feedback-send-button]').should('not.be.disabled');
  });
});
