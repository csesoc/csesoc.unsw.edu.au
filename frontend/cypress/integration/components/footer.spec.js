describe("Footer Testing", () => {
    beforeEach(() => {
        cy.visit('/');
    });

    it('finds whether the footer links and the content is displayed', () => {
        // Check if the logo has link
        cy.get('[data-cy=footer-logo]');

        // Check if the address is displayed
        cy.get('[data-cy=footer-address]').contains('B03 CSE Building K17, UNSW');

        // Check all the internal link is present
        cy.get('[data-cy=footer-internal-link] > a').should("have.length", 5).each(($link) => {
            expect($link).to.have.attr('href');
        });

        // Check all the media link is present
        cy.get('[data-cy=footer-media-link] > a').should("have.length", 8).each(($link) => {
            expect($link).to.have.attr('href');
        });
        // Check the media title is display
        cy.get('[data-cy=footer-media-title]').contains('Social Media');

        // Check all the resources link is present
        cy.get('[data-cy=footer-resources-link] > a').should("have.length", 5).each(($link) => {
            expect($link).to.have.attr('href');
        });
        // Check the resources title is display
        cy.get('[data-cy=footer-resources-title]').contains('For your better future');

 
    });
})