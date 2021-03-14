describe('Mission Landing Page', () => {
  beforeEach(() => {
    cy.visit('/');
  });

  it('finds whether statement and graphic are displayed', () => {
    //  Make sure it is visible.
    cy.get('[data-cy=mission-img]').scrollIntoView().should('be.visible');
    // Make sure that there is a statement, whatever that statement may be.
    cy.get('[data-cy=mission-statement]').contains(/\w+/);
  });
});