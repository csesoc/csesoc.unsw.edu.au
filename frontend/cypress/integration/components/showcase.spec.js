describe('Showcase Landing Page', () => {
    beforeEach(() => {
      cy.visit('/');
    });
    
    it('tests scrolling to the join us section from the top of the landing page', () => {
      // finding the join us button and clicking it

      // cypress is giving a 'the chainer inViewPort cannot be found', 
      // might try making it a command not an assert to see if that fixes it
      // ref: https://github.com/cypress-io/cypress/issues/877
      // cy.get('[data-cy=communityLinks]').should('not.be.inViewPort');
      // cy.get('[data-cy=joinus-button]').click();
      // it should scroll to the join us section
      // cy.get('[data-cy=communityLinks]').should('not.be.inViewPort');
      
    });
    
  });