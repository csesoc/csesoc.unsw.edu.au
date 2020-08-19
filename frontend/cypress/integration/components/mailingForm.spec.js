describe('Mailing forms', () => {
  it('checks labels and validation of a general form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // Check if name label exists
    cy.get('[data-cy=name-label]').contains('Name *');
    // Check if email label exists
    cy.get('[data-cy=email-label]').contains('Email *');
    // Check if message label exists
    cy.get('[data-cy=message-label]').contains('Message *');
    // Ensure send button is disabled
    cy.get('[data-cy=send-button]').should('be.disabled');
    // Fill in valid name
    cy.get('[data-cy=name-field]').type('Sergio');
    // Fill in valid email
    cy.get('[data-cy=email-field]').type('sergio@gmail.com');
    // Fill in valid message
    cy.get('[data-cy=message-field]').type('message goes here');
    // Ensure send button is enabled after entering valid inputs
    cy.get('[data-cy=send-button]').should('not.be.disabled');
  });

  it('checks labels and validation of a sponsorship form', () => {
    // Visit sponsors page
    cy.visit('/#/sponsors');
    // Check if name label exists
    cy.get('[data-cy=name-label]').contains('Company Name *');
    // Check if email label exists
    cy.get('[data-cy=email-label]').contains('Email *');
    // Check if message label exists
    cy.get('[data-cy=message-label]').contains('Message *');
    // Ensure send button is disabled
    cy.get('[data-cy=send-button]').should('be.disabled');
    // Fill in valid name
    cy.get('[data-cy=name-field]').type('Sergio');
    // Fill in invalid email
    cy.get('[data-cy=email-field]').type('sergio2@ema.il.c');
    // Fill in valid message
    cy.get('[data-cy=message-field]').type('message goes here');
    // Ensure send button is disabled,
    // because the email is invalid
    cy.get('[data-cy=send-button]').should('be.disabled');
    // Fill in valid email
    cy.get('[data-cy=email-field]').clear().type('sergio@gmail.com');
    // Ensure send button is enabled after making the email valid
    cy.get('[data-cy=send-button]').should('not.be.disabled');
  });

  it('checks labels and validation of a feedback form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // Check if name label exists
    cy.get('[data-cy=name-label]').contains('Name');
    // Check if email label exists
    cy.get('[data-cy=email-label]').contains('Email');
    // Check if message label exists
    cy.get('[data-cy=message-label]').contains('Message *');
    // Ensure send button is disabled
    cy.get('[data-cy=send-button]').should('be.disabled');
    // Fill in valid message
    cy.get('[data-cy=message-field]').type('message goes here');
    // Ensure send button is enabled,
    // because the message is the only required field
    cy.get('[data-cy=send-button]').should('not.be.disabled');
  });
});
