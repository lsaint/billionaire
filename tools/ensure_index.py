# -*- coding: utf-8 -*-

from pymongo import MongoClient

db = MongoClient("localhost", 27017).billionaire
name = raw_input("name: ")
pw   = raw_input("pw: ")
db.authenticate(name, pw)

print db.gift.ensure_index("uid")

print db.sponsor.ensure_index("uid")
