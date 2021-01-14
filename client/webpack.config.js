const { resolve, join } = require("path");
const webpack = require("webpack");
const autoPrefixer = require('autoprefixer');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const publicFolder = resolve("./dist");
const { ENV } = process.env;

const isProd = ENV === "production";

const webpackLoader = {
  loader: "elm-webpack-loader",
  options: {
    debug: isProd ? false : true,
    optimize: isProd,
    cwd: __dirname,
  },
};

module.exports = {
  mode: isProd ? "production" : "development",
  entry: "./src/index.js",
  devServer: {
    contentBase: publicFolder,
    hot: true,
    publicPath: '/',
    port: 3000,
    watchContentBase: true
  },
  output: {
    filename: 'bundle.js',
    publicPath: '/',
    path: publicFolder
  },
  module: {
    rules: [
      {
        test: /(\.css)$/,
        use: ['style-loader', 'css-loader']
      }, {
        test: /\.sc?ss$/,
        use: [
          'style-loader',
          {
            loader: 'css-loader',
            options: { importLoaders: 2 }
          },
          {
            loader: 'postcss-loader',
            options: {
              postcssOptions: {
                plugins: (loader) => [
                  require('postcss-import')(),
                  autoPrefixer(),
                ]
              }
            }
          },
          { loader: 'sass-loader' }
        ],
      }, {
        test: /\.elm$/,
        exclude: [/elm-stuff/, /node_modules/],
        use: isProd
          ? [webpackLoader]
          : [{ loader: "elm-hot-webpack-loader" }, webpackLoader],
      },
      {
        test: /\.js$/,
        enforce: 'pre',
        use: ['source-map-loader'],
      },
    ],
  },
  plugins: [
    new webpack.NoEmitOnErrorsPlugin(),
    new webpack.LoaderOptionsPlugin({
      options: {
        postcss: [autoPrefixer()]
      }
    }),
    new CopyWebpackPlugin([
      { from: 'assets/images', to: 'images' }
    ]),
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
      favicon: 'assets/favicon.ico'
    })
  ],
};