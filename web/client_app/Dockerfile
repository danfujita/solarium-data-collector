FROM node:8

EXPOSE 3000
RUN npm install webpack -g

RUN mkdir /frontend

WORKDIR /frontend
COPY package.json /frontend/
RUN npm config set registry http://registry.npmjs.org/ && npm install

ADD . /frontend/
CMD  npm run dev
