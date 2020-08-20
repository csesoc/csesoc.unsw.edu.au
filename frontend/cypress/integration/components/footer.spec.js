describe("Footer Testing", () => {
    beforeEach(() => {
        cy.visit('/');
    });

    it('finds whether the logo links to the homepage and the address is displayed', () => {
        // Check if the logo has link
        cy.get('[data-cy=footer-main-logo]').should('have.attr', 'href', 'localhost:8080/');

        // Check if the address is displayed
        cy.get('[data-cy=footer-address]').contains('B03 CSE Building K17, UNSW');

        // Check if the internal link is displayed

        // cy.get('[data-cy=footer-internal-link]').should( ($lis) => {
            // expect($lis, '4 items').to.have.length(4);
            // expect($lis, '')
        // })
    });
})