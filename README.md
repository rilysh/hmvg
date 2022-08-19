## How-many-views-github
Mini HTTP server, counts GitHub profile views

<p align="center">
<img src="https://user-images.githubusercontent.com/71683721/185570540-f06ccc5d-bc9e-4921-b514-3c693e7972b2.png"/>
</p>

## Usage
- Go to [github.com](https://github.com/) and create a repo that is the same as your username
- Add a README file to that repo
- Start editing the file and add this URL mentioned below
```
![](https://hmvg.herokuapp.com?username=your_github_username)
```
- Now check preview, you should have an image with view counts (if you haven't used this before, the first count will be 0 and increase on each view)

💡 To change colors, try these parameters

### To change "Profile views" background color
![1](https://user-images.githubusercontent.com/71683721/185572004-97ac6138-7c59-4602-aeeb-574ace450a5f.png)

`![](https://hmvg.herokuapp.com?username=your_github_username&first_color=b54ed4)`

### To change counts background color
![2](https://user-images.githubusercontent.com/71683721/185572057-31f6b172-dde4-464a-b913-3f86a747dfb8.png)

`![](https://hmvg.herokuapp.com?username=your_github_username&second_color=d4514e)`
 
### If you want to change both sides color
![3](https://user-images.githubusercontent.com/71683721/185572086-f432d7b0-8af5-4cd8-9055-134522651733.png)

`![](https://hmvg.herokuapp.com?username=your_github_username&first_color=d4514e&second_color=de7e1f)`

## Installation
#### To run it locally
- Clone the repo
`git clone https://github.com/kiwimoe/hmvg.git`

- Move there and run
`go build`*

- Now run the server, should be active at `localhost:1337`

Note: "*" You must have to install Go before taking furthur steps. To download visit [go.dev](https://go.dev/dl/)

#### To deploy on heroku
To learn how to deploy your own instance, visit Heroku [docs](https://devcenter.heroku.com/articles/git)
