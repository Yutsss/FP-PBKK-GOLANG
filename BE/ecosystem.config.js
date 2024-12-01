module.exports = {
    apps: [
        {
            name: 'fp_pbkk_go',
            script: './backend',
            instances: '1',
            env: {
                GO_ENV: 'development',
            },
            env_production: {
                GO_ENV: 'production',
            }
        }
    ]
};
