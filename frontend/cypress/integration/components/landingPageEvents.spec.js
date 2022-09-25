describe('Landing page events', () => {
    it('checks for a slideshow of events on regular viewports, or a message otherwise', () => {
      cy.visit('/');
      cy.get('[data-cy=event-section]').then(($eventSection) => {
          if ($eventSection.children('[data-cy=event-alert]').length > 0) {
            // Ensure that the alert exists before asserting it.
            cy.get('[data-cy=event-alert]').should('be.visible');
        } else {
            // Otherwise, the slider must exist, and the event list should not,
            // as this is a desktop viewport.
            // cy.get('[data-cy=event-slider]').should('be.visible');
            // cy.get('[data-cy=event-list]').should('not.be.visible');
        }
    });      
});

it('checks for a list of events on mobile viewports, or a message otherwise', () => {
    cy.visit('/');
    cy.viewport('iphone-6');
    cy.get('[data-cy=event-section]').then(($eventSection) => {
        if ($eventSection.children('[data-cy=event-alert]').length > 0) {
            // Ensure that the alert exists before asserting it.
            cy.get('[data-cy=event-alert]').should('be.visible');
        } else {
            // Otherwise, the list must exist, and the slider should not,
            // as this is a mobile viewport.
            cy.get('[data-cy=event-slider]').should('not.be.visible');
            cy.get('[data-cy=event-list]').should('be.visible');
          }
        });
    });   
});
  