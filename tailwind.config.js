/** @type {import('tailwindcss').Config} */
    module.exports = {
        content: [
            "./views/**/*.{html,tmpl,js}",
            "./public/scripts/**/*.js"

        ],
        theme: {
            extend: {
                keyframes: {
                    slideInLeft: {
                        '0%': { opacity: 0, transform: 'translateX(-20px)' },
                        '100%': { opacity: 1, transform: 'translateX(0)' },
                    },
                    slideInRight: {
                        '0%': { opacity: 0, transform: 'translateX(20px)' },
                        '100%': { opacity: 1, transform: 'translateX(0)' },
                    },
                    fadeInScale: {
                        '0%': { opacity: 0, transform: 'translateY(10px) scale(0.95)' },
                        '100%': { opacity: 1, transform: 'translateY(0) scale(1)' },
                    },
                    fadeInLeft: {
                        '0%': { opacity: 0, transform: 'translateX(-20px)' },
                        '100%': { opacity: 1, transform: 'translateX(0)' },
                    },
                    fadeInRight: {
                        '0%': { opacity: 0, transform: 'translateX(20px)' },
                        '100%': { opacity: 1, transform: 'translateX(0)' },
                    },
                },
                animation: {
                    'fade-in-left': 'fadeInLeft 0.4s ease-out forwards',
                    'fade-in-right': 'fadeInRight 0.4s ease-out forwards',
                    'slide-in-left': 'slideInLeft 0.4s ease-out forwards',
                    'slide-in-right': 'slideInRight 0.4s ease-out forwards',
                    'fade-in-scale': 'fadeInScale 0.4s ease-out forwards',
                },
            },
        },
        plugins: [],
    }

