describe('Mailing forms', () => {
  it('checks labels and validation of a general form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // Check if name label exists
    cy.get('[data-cy=mailing-name]').contains('Name *');
    // Check if email label exists
    cy.get('[data-cy=mailing-email]').contains('Email *');
    // Check if message label exists
    cy.get('[data-cy=mailing-message]').contains('Message *');
  });

  it('checks labels and validation of a sponsorship form', () => {
    // Visit sponsors page
    cy.visit('/#/sponsors');
    // Check if name label exists
    cy.get('[data-cy=mailing-name]').contains('Company Name *');
    // Check if email label exists
    cy.get('[data-cy=mailing-email]').contains('Email *');
    // Check if message label exists
    cy.get('[data-cy=mailing-message]').contains('Message *');
  });

  it('checks labels and validation of a feedback form', () => {
    // Visit engage page
    cy.visit('/#/engage');
    // Check if name label exists
    cy.get('[data-cy=mailing-name]').contains('Name');
    // Check if email label exists
    cy.get('[data-cy=mailing-email]').contains('Email');
    // Check if message label exists
    cy.get('[data-cy=mailing-message]').contains('Message *');
  });
});
