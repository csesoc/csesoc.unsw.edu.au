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

        cy.get('[data-cy=footer-internal-link]').should( ($lis) => {
            expect($lis, '4 items').to.have.length(4);
            expect($lis.eq(0), 'first item').to.have.attr('href', 'localhost:8080/about');
            expect($lis.eq(0), 'first item').to.have.text('About');
            expect($lis.eq(1), 'second item').to.have.attr('href', 'localhost:8080/resources');
            expect($lis.eq(1), 'second item').to.have.text('Resources');
            expect($lis.eq(2), 'third item').to.have.attr('href', 'localhost:8080/sponsor');
            expect($lis.eq(2), 'third item').to.have.text('Sponsor');
            expect($lis.eq(3), 'forth item').to.have.attr('href', 'localhost:8080/engage');
            expect($lis.eq(3), 'first item').to.have.text('Engage');
        });
    });
})