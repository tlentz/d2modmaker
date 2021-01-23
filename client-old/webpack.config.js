const path = require("path");
const { merge } = require('webpack-merge');

const CopyWebpackPlugin = require("copy-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");

// Production CSS assets - separate, minimised file
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");

const autoPrefixer = require('autoprefixer');

var MODE =
    process.env.npm_lifecycle_event === "prod" ? "production" : "development";
var withDebug = !process.env["npm_config_nodebug"] && MODE === "development";
console.log(
    "\x1b[36m%s\x1b[0m",
    `** elm-webpack-starter: mode "${MODE}", withDebug: ${withDebug}\n`
);

var common = {
    mode: MODE,
    entry: "./src/index.js",
    output: {
        path: path.join(__dirname, "dist"),
        publicPath: "/",
        // FIXME webpack -p automatically adds hash when building for production
        filename: MODE === "production" ? "[name]-[hash].js" : "index.js"
    },
    plugins: [
        new HtmlWebpackPlugin({
            inject: 'body',
            filename: 'index.html',
            template: require('html-webpack-template'),
            appMountId: 'main',
            mobile: true,
            lang: 'en-US',
            title: 'Elm-D2MM',
            links: [],
            xhtml: true,
            hash: false,
            chunks: ['main'],
            favicon: 'src/assets/images/logo.png'
        })
    ],
    resolve: {
        modules: [path.join(__dirname, "src"), "node_modules"],
        extensions: [".js", ".elm", ".scss", ".png"]
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.css$/,
                exclude: [/elm-stuff/, /node_modules/],
                use: ["style-loader", "css-loader?url=false", "postcss-loader"]
            },
            {
                test: /\.css$/,
                exclude: [/elm-stuff/, /node_modules/],
                use: ["style-loader", "css-loader?url=false"]
            },
            {
                test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
                exclude: [/elm-stuff/, /node_modules/],
                use: {
                    loader: "url-loader",
                    options: {
                        limit: 10000,
                        mimetype: "application/font-woff"
                    }
                }
            },
            {
                test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
                exclude: [/elm-stuff/, /node_modules/],
                loader: "file-loader"
            },
            {
                test: /\.(jpe?g|png|gif|svg)$/i,
                exclude: [/elm-stuff/, /node_modules/],
                loader: "file-loader"
            }
        ]
    }
};

if (MODE === "development") {
    module.exports = merge(common, {
        optimization: {
            // Prevents compilation errors causing the hot loader to lose state
            noEmitOnErrors: true
        },
        module: {
            rules: [
                {
                    test: /\.elm$/,
                    exclude: [/elm-stuff/, /node_modules/],
                    use: [
                        { loader: "elm-hot-webpack-loader" },
                        {
                            loader: "elm-webpack-loader",
                            options: {
                                // add Elm's debug overlay to output
                                debug: withDebug,
                            }
                        }
                    ]
                }
            ]
        },
        devServer: {
            inline: true,
            stats: "errors-only",
            contentBase: path.join(__dirname, "src/assets"),
            historyApiFallback: true,
            // feel free to delete this section if you don't need anything like this
            before(app) {
                // on port 3000
                app.get("/test", function (req, res) {
                    res.json({ result: "OK" });
                });
            }
        }
    });
}

if (MODE === "production") {
    module.exports = merge(common, {
        optimization: {
            minimizer: [
                new OptimizeCSSAssetsPlugin({})
            ]
        },
        plugins: [
            // Delete everything from output-path (/dist) and report to user
            new CleanWebpackPlugin({
                root: __dirname,
                exclude: [],
                verbose: true,
                dry: false
            }),
            // Copy static assets
            new CopyWebpackPlugin({
                patterns: [
                    {
                        from: "src/assets"
                    }
                ]
            }),
            new MiniCssExtractPlugin({
                // Options similar to the same options in webpackOptions.output
                // both options are optional
                filename: "[name]-[hash].css"
            })
        ],
        module: {
            rules: [
                {
                    test: /\.elm$/,
                    exclude: [/elm-stuff/, /node_modules/],
                    use: {
                        loader: "elm-webpack-loader",
                        options: {
                            optimize: true
                        }
                    }
                }
            ]
        }
    });
}
