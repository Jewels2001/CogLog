## <img src="https://github.com/Jewels2001/CogLog/assets/53021785/3716b232-c60c-466d-897a-01bca1d325b4" width="300">
_A brain-to-data platform by_  <img src="https://github.com/Jewels2001/CogLog/assets/53021785/1479f025-a925-41cd-b8c3-2952072c47fb" width="60">.

For [natHACKS 2023](https://nathacks23.devpost.com/) focused on brain-to-text and speech-to-text capabilities.

///

## Intro
Brain-Computer Interface (BCI) equipment allows people to connect with their environment and other people with greater awareness and capability. CogLog has produced a SAAS called Cerecode that leverages BCI technology to convert brain wave readings into spoken speech. This gives people who have limited or decreasing speech capabilities a way to communicate more clearly with those around them. More specifically, Cerecode utilizes models from the open-source [dataset](https://datadryad.org/stash/dataset/doi:10.5061/dryad.x69p8czpq) associated with [_"A high-performace speech neruprosthesis"_](https://www.nature.com/articles/s41586-023-06377-x.pdf) which focuses on decoding phonemes from intracortical microelectrode arrays. 

## Usage
Client-Side:

`npm install`

`npm run dev`

Decoder Model:
- Follow instructions from [speechBCI](https://github.com/fwillett/speechBCI/tree/main) repo.
> [!IMPORTANT]
> We used a 16 vCPU + 128GB RAM VM to host the model. 64GB was not enough.

Gateway Service:

`go run server.go`

## Roadmap for Future
We can improve upon our SAAS by providing a fully functioning platform, and implementing tools such as wearables and mobile apps. This could also lead to the ability to make a natural recording of one's thoughts, such as live note-taking (during lectues or meetings) or an idea journal.
We could also personalize this usecase, by providing support for multiple voices and enable real time decoding of brain data.
Further, we could improve upon our decoding performace, possibly writing our own BCI decoder model fromm scratch with a larger and more robust dataset.


## Contributors
- Ali Amusat: Neuroscience Specialist, Graphic and UI/UX Designer
- Rozella Ricci Dagta: Graphic and UI/UX Designer, Full Stack
- Alex Hughes: Business Strategist, Frontend
- Dustin Ward: Backend, Full Stack
- Julie Wojtiw-Quo: ML Specialist, Full Stack

## Project Links
- [Devpost](https://devpost.com/software/coglog)
- [Slide Deck](https://docs.google.com/presentation/d/1Xrh-y_W0OSOFUX3jnuKIfwNEUtSFSuoYoYQwCp_6Km8/edit?usp=sharing)
- [Video](https://youtu.be/DVbrny5-V0c)

## Resources
- [Dataset](https://datadryad.org/stash/dataset/doi:10.5061/dryad.x69p8czpq)
- [Dataset's associated codebase](https://github.com/fwillett/speechBCI/tree/main)
- [KARA One study and dataset](http://www.cs.toronto.edu/~complingweb/data/karaOne/karaOne.html)
