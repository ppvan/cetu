/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./ui/html/**/*.html",
        "./ui/static/js/**/*.js",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ]
}

